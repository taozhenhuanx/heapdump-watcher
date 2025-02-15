package cli

import (
	"fmt"
	"heapdump_watcher/controller/store"
	"heapdump_watcher/controller/store/aliyun"
	"heapdump_watcher/setting"
)

// 文件上传
func UPload(uploadFile, filePath string) (error, string) {
	var (
		uploader store.Uploader
		err      error
	)
	// 判断类型
	switch setting.Conf.StorageInfo.StorageType {
	case "aliyun":
		uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
			// 先写死,也可以从命令行传递
			EndPoint:     setting.Conf.StorageInfo.OssEndpoint,
			AccessKey:    setting.Conf.StorageInfo.AccessKey,
			AccessSecret: setting.Conf.StorageInfo.AccessSecret,
		})
	case "tx":
	case "aws":
	default:
		return fmt.Errorf("不支持该厂商类型"), ""
	}
	if err != nil {
		return err, ""
	}

	// 如果上面校验密码正确,那么就使用upload来上传文件
	return uploader.Upload(setting.Conf.StorageInfo.BucketName, uploadFile, filePath)
}
