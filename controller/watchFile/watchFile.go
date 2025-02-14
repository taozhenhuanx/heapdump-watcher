package watchfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"heapdump_watcher/controller/sendAlert"
	"heapdump_watcher/controller/store/cli"

	"github.com/fsnotify/fsnotify"
)

// event.Name 是当前正在被监听的文件路径+文件名
func WatchFiles() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create fsnotify watcher: %v", err)
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				// 检测到新文件创建 (strings.HasSuffix函数检查prof后缀)
				if strings.HasSuffix(event.Name, "prof") {
					log.Printf("New heapdump file detected: %s", event.Name)
					// 等待文件写入完成
					if ok := isFileComplete(event.Name, 30*time.Second, 2*time.Second); !ok {
						log.Printf("Failed to wait for file completion: %v", err)
						continue
					}
					// // 上传文件到OSS,  appName OSS的URL
					appName := filepath.Base(filepath.Dir(filepath.Dir(event.Name)))
					err := cli.UPload(event.Name, appName)
					if err != nil {
						log.Printf("Failed to upload file to OSS: %v", err)
					} else {
						log.Printf("File uploaded to OSS successfully: %s", event.Name)
						// 发送告警通知
						sendAlert.SendAlertType(appName)
						if err != nil {
							log.Printf("Failed to send WeChat alert: %v", err)
						}
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Error: %v", err)
		}
	}
}

// 判断文件是否写入完成
// isFileComplete 检查文件的大小是否在指定的最大时间内不变
func isFileComplete(filePath string, maxDuration, checkInterval time.Duration) bool {
	var initialSize int64
	var err error

	// 获取初始文件大小
	if initialSize, err = getFileSize(filePath); err != nil {
		log.Println("Error getting file size:", err)
		return false
	}

	// 开始检查文件大小
	startTime := time.Now()
	// time.Since(startTime)  计算startTime 到现在的时间，然后判断是不是超过了我们要比较的时间maxDuration
	for time.Since(startTime) < maxDuration {
		time.Sleep(checkInterval) // 等待检查间隔

		finalSize, err := getFileSize(filePath)
		if err != nil {
			log.Println("Error getting file size:", err)
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
		log.Println("文件不存在:", err)
		return 0, nil // 文件不存在
	}
	if err != nil {
		return 0, err
	}
	fmt.Println("文件大小", fileInfo.Size())
	return fileInfo.Size(), nil
}
