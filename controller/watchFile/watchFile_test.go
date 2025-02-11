package watchfile

import (
	"os"
	"testing"
	"time"
)

// TestIsFileComplete 测试 isFileComplete 函数
func TestIsFileComplete(t *testing.T) {
	// 检查文件大小是否在 30 秒内不变，每 2 秒检查一次
	if isStable := isFileComplete("filepath", 30*time.Second, 2*time.Second); isStable {
		fmt.Println("File size is stable.")
	} else {
		fmt.Println("File size changed during the check.")
	}
}
