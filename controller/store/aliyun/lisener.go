package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

// 构造函数
func NewDefaultProgressListener() *ProgressListener {
	return &ProgressListener{}
}

// 显示进度条progressbar
type ProgressListener struct {
	bar *progressbar.ProgressBar
}

// 显示上传到Oss显示上传进度用的
func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		p.bar = progressbar.DefaultBytes(
			event.TotalBytes,
			"文件上传中",
		)
	case oss.TransferDataEvent:
		// Add64()上传总大小
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\n上传完成\n")
	case oss.TransferFailedEvent:
		fmt.Printf("\n上传失败\n")
	default:
	}
}
