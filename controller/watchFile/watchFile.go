package watchFile

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"heapdump_watcher/controller/k8sUtils"
	"heapdump_watcher/controller/sendAlert"
	"heapdump_watcher/controller/store/cli"
	"heapdump_watcher/setting"
	"heapdump_watcher/utils"

	"github.com/fsnotify/fsnotify"
)

// event.Name 是当前正在被监听的文件路径+文件名
func WatchFiles(watcher *fsnotify.Watcher) {
	// 添加监听目录
	if err := watcher.Add(setting.Conf.FilePath.WatchPath); err != nil {
		logrus.Fatal("Add failed:", err)
	}
	// done := make(chan bool)
	// 开始监听事件
	go func() {
		// defer close(done)
		for {
			select {
			// watcher.Events 事件
			case event, ok := <-watcher.Events:
				if !ok {
					logrus.Errorf("watcher Events 通道已关闭")
					return
				}
				// 判断监听事件是Create
				if event.Op&fsnotify.Create == fsnotify.Create {
					// 检测到新文件创建 (strings.HasSuffix函数检查prof后缀)
					if strings.HasSuffix(event.Name, "prof") {
						logrus.Printf("检测到新的heap dump文件: %s", event.Name)

						// 等待文件写入完成
						if ok := isFileComplete(event.Name, 30*time.Second, 2*time.Second); !ok {
							logrus.Errorf("等待文件完成失败: %v", ok)
							continue
						}

						// 压缩文件
						zipFilePath, err := utils.ZipFile(event.Name)
						if err != nil {
							logrus.Errorf("压缩文件失败: %v", err)
							continue
						}

						//  appName OSS的URL  [filepath.Dir 获取目录]
						appName := filepath.Base(zipFilePath)
						podName, err := utils.GetFileNameWithoutExt(event.Name)
						if err != nil {
							logrus.Errorf("GetFileNameWithoutExt: %v", err)
							continue
						}
						// k8s cient-go
						clientset, err := setting.ReadKubeConf()
						if err != nil {
							logrus.Errorf("ReadKubeConf Error: %s", err)
						}

						// 获取名称空间名字
						nsName, err := k8sUtils.GetPodNamespace(clientset, podName)
						if err != nil {
							logrus.Errorf("获取名称空间名字 Error: %s", err)
						}

						// 上传文件到OSS
						// event.Name 监听的文件绝对路径
						err, ossURL := cli.UPload(appName, zipFilePath)
						if err != nil {
							logrus.Errorf("Failed to upload file to OSS: %v", err)
							continue
						}

						// 发送告警通知 ossURL, podName, nsName
						if err := sendAlert.SendAlertType(ossURL, podName, nsName); err != nil {
							logrus.Errorf("发送告警失败: %s", err)
							continue
						}

						// 清理生成的 ZIP 文件
						if err := os.Remove(zipFilePath); err != nil {
							logrus.Errorf("清理生成的 ZIP 文件: %s", err)
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

	logrus.Println("heapdump-watcher 程序已经启动")
	// 永久阻塞 main goroutine
	// <-done
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
	logrus.Printf("等待文件生成完成，当前文件大小是: %dM", fileInfo.Size()/1048576)
	return fileInfo.Size(), nil
}
