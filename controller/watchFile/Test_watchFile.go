package watchfile

import (
	"os"
	"testing"
	"time"
)

// TestIsFileComplete 测试 isFileComplete 函数
func TestIsFileComplete(t *testing.T) {
	// 创建一个临时文件
	tempFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // 清理临时文件

	// 初始文件大小
	initialSize := int64(0)
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// 检查文件写入完成的逻辑
	if complete := isFileComplete(tempFile.Name(), 30*time.Second, 2*time.Second); !complete {
		t.Error("Expected file to be complete, but it was not.")
	}

	// 模拟文件大小变化
	if err := os.Truncate(tempFile.Name(), 0); err != nil {
		t.Fatalf("Failed to truncate temp file: %v", err)
	}

	// 由于我们已经修改了文件，我们需要再次检查
	if complete := isFileComplete(tempFile.Name(), 30*time.Second, 2*time.Second); complete {
		t.Error("Expected file to not be complete after truncation, but it was.")
	}

	// 恢复文件内容
	if err := os.WriteFile(tempFile.Name(), []byte("final data"), 0644); err != nil {
		t.Fatalf("Failed to restore temp file: %v", err)
	}

	// 再次检查文件写入完成的逻辑
	if complete := isFileComplete(tempFile.Name(), 30*time.Second, 2*time.Second); !complete {
		t.Error("Expected file to be complete after restoring, but it was not.")
	}
}
