package utils

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// TestZipFile 测试 ZipFile 函数
func TestZipFile(t *testing.T) {
	// 创建一个临时文件
	tempFile, err := ioutil.TempFile("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %s", err)
	}
	defer os.Remove(tempFile.Name()) // 清理临时文件

	// 向临时文件写入一些内容
	content := []byte("Hello, World!")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("写入临时文件失败: %s", err)
	}
	tempFile.Close() // 关闭文件，以便后续操作

	// 调用 ZipFile 函数
	zipPath, err := ZipFile(tempFile.Name())
	if err != nil {
		t.Fatalf("压缩失败: %s", err)
	}
	defer os.Remove(zipPath) // 清理生成的 ZIP 文件

	// 检查 ZIP 文件是否存在
	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		t.Fatalf("ZIP 文件不存在: %s", zipPath)
	}

	// 解压缩 ZIP 文件以验证内容
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		t.Fatalf("打开 ZIP 文件失败: %s", err)
	}
	defer r.Close()

	// 检查 ZIP 文件中的所有条目
	if len(r.File) != 1 {
		t.Fatalf("ZIP 文件应该包含一个条目")
	}

	// 提取第一个条目
	zippedFile := r.File[0]
	rc, err := zippedFile.Open()
	if err != nil {
		t.Fatalf("打开 ZIP 条目失败: %s", err)
	}
	defer rc.Close()

	// 读取 ZIP 文件中的内容
	unzippedContent, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Fatalf("读取 ZIP 条目失败: %s", err)
	}

	// 验证内容是否正确
	if string(unzippedContent) != "Hello, World!" {
		t.Errorf("解压缩后的内容不正确: %s", string(unzippedContent))
	}
}

func TestGetFileNameWithoutExt(t *testing.T) {
	fileName, _ := GetFileNameWithoutExt("/app/xxx/qqq/qi-capability-68896b6858-zwmfb.hprof")
	fmt.Println(fileName)
}
