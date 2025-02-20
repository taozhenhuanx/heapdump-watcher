package utils

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// ZipFile 接收文件路径，压缩该文件并返回压缩文件的路径
func ZipFile(filePath string) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Format("20060102150405") // 格式为 YYYYMMDDHHMMSS

	// 从传入的文件路径中提取文件名和目录
	dir := filepath.Dir(filePath)                      // 获取文件所在目录
	fileName := filepath.Base(filePath)                // 获取文件名
	zipFileName := fileName + "-" + timestamp + ".zip" // 生成压缩文件名
	zipFilePath := filepath.Join(dir, zipFileName)     // 压缩文件的完整路径

	// 创建一个新的 ZIP 文件
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		logrus.Errorf("无法创建 ZIP 文件: %s", err)
		return "", err
	}
	defer zipFile.Close() // 确保函数结束时关闭文件

	// 创建新的 ZIP 写入器
	writer := zip.NewWriter(zipFile)
	defer writer.Close() // 确保写入器关闭

	// 打开要压缩的文件
	sourceFile, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("无法打开源文件: %s", err)
		return "", err
	}
	defer sourceFile.Close() // 确保源文件关闭

	// 创建 ZIP 文件中的文件
	zipEntryWriter, err := writer.Create(fileName) // 使用原文件名作为 ZIP 条目名
	if err != nil {
		logrus.Errorf("无法创建 ZIP 条目: %s", err)
		return "", err
	}

	// 将源文件内容写入 ZIP 文件
	_, err = io.Copy(zipEntryWriter, sourceFile)
	if err != nil {
		logrus.Errorf("无法写入 ZIP 文件: %s", err)
		return "", err
	}
	logrus.Info("写入 ZIP 文件完成")
	// 返回压缩后的文件路径
	return zipFilePath, nil
}

// 传入文件全路径, 返回文件名字
func GetFileNameWithoutExt(filePath string) (string, error) {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err // 文件不存在，返回空字符串
	}
	fileName := path.Base(filePath)
	// 去除扩展名
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return fileNameWithoutExt, nil
}
