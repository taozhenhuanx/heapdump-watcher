package watchFile

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"heapdump_watcher/controller/sendAlert"
	"heapdump_watcher/controller/store/cli"
	"heapdump_watcher/setting"

	"github.com/fsnotify/fsnotify"
)

// event.Name 是当前正在被监听的文件路径+文件名
func WatchFiles() {
	// 创建一个监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.Fatalf("Failed to create fsnotify watcher: %v", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	// 开始监听事件
	go func() {
		defer close(done)
		for {
			select {
			// watcher.Events 事件
			case event, ok := <-watcher.Events:
				if !ok {
					logrus.Fatalf("watcher Events Error")
					return
				}
				// 判断监听事件是Create
				if event.Op&fsnotify.Create == fsnotify.Create {
					// 检测到新文件创建 (strings.HasSuffix函数检查prof后缀)
					if strings.HasSuffix(event.Name, "prof") {
						logrus.Printf("检测到新的heap dump文件: %s", event.Name)

						// 等待文件写入完成
						if ok := isFileComplete(event.Name, 60*time.Second, 2*time.Second); !ok {
							logrus.Printf("等待文件完成失败: %v", err)
							continue
						}
						// if zipFilePath, err := utils.ZipFile(event.Name); err != nil {
						// 	fmt.Println(zipFilePath)
						// }
						// 上传文件到OSS,  appName OSS的URL  [filepath.Dir 获取目录、]
						appName := filepath.Base(filepath.Dir(filepath.Dir(event.Name)))
						logrus.Println("appName 是", appName)
						err, OssURL := cli.UPload(appName, event.Name)
						if err != nil {
							logrus.Printf("Failed to upload file to OSS: %v", err)
							continue
						}

						// 发送告警通知
						if err := sendAlert.SendAlertType(OssURL); err != nil {
							logrus.Printf("发送告警失败: %s", err)
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logrus.Printf("Error: %v", err)
			}
		}
	}()

	// 添加监听目录
	if err := watcher.Add(setting.Conf.FilePath.WatchPath); err != nil {
		logrus.Fatal("Add failed:", err)
	}
	// 永久阻塞 main goroutine
	<-done
}

// 判断文件是否写入完成
// isFileComplete 检查文件的大小是否在指定的最大时间内不变
func isFileComplete(filePath string, maxDuration, checkInterval time.Duration) bool {
	var initialSize int64
	var err error

	// 获取初始文件大小
	if initialSize, err = getFileSize(filePath); err != nil {
		logrus.Error("Error getting file size:", err)
		return false
	}

	// 开始检查文件大小
	startTime := time.Now()
	// time.Since(startTime)  计算startTime 到现在的时间，然后判断是不是超过了我们要比较的时间maxDuration
	for time.Since(startTime) < maxDuration {
		time.Sleep(checkInterval) // 等待检查间隔

		finalSize, err := getFileSize(filePath)
		if err != nil {
			logrus.Error("Error getting file size:", err)
			return false
		}

		if finalSize != initialSize {
			// 如果文件大小变化，重置初始大小并重置计时
			initialSize = finalSize
			startTime = time.Now()
		}
	}

	return true
}

// 获取文件大小
func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		logrus.Error("文件不存在:", err)
		return 0, nil // 文件不存在
	}
	if err != nil {
		return 0, err
	}
	logrus.Info("文件大小", fileInfo.Size())
	return fileInfo.Size(), nil
}
