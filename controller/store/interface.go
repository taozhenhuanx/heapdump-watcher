package store

// 定义上传文件到bucket
// 做了抽象,并不关心我们需要上传到哪个Oss厂商的bucket
// filePath    the local file path to upload.
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) (error, string)
}
