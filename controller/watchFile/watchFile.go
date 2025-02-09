package watchfile

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watchFiles() {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				// 检测到新文件创建
				if strings.HasSuffix(event.Name, "heapdump.prof") {
					log.Printf("New heapdump file detected: %s", event.Name)
					// 等待文件写入完成
					if err := waitForFileCompletion(event.Name); err != nil {
						log.Printf("Failed to wait for file completion: %v", err)
						continue
					}
					// 上传文件到OSS
					appName := filepath.Base(filepath.Dir(filepath.Dir(event.Name)))
					err := uploadFileToOSS(event.Name, appName)
					if err != nil {
						log.Printf("Failed to upload file to OSS: %v", err)
					} else {
						log.Printf("File uploaded to OSS successfully: %s", event.Name)
						// 发送企业微信告警通知
						err = sendWechatAlert(appName)
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
