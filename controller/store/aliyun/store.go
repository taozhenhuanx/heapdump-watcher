package aliyun

import (
	"fmt"
	"heapdump_watcher/controller/store"
	"heapdump_watcher/setting"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
判断对象是否实现了接口的强制约束   【也就是保证了对象实现的方法一定要跟抽象的接口的方法保持一致】
a string = "abc"
_ store.Uploader 我不需要这个变量的值, 我只是做变量类型的判断
&AliOssStore{} 这个对象 必须满足 store.Uploader
_ store.Uploader = &AliOssStore{} 声明了一个空对象, 只是需要一个地址
nil 空指针, nil有没哟类型: 有类型
a *AliOssStore = nil   nil是一个AliOssStore 的指针
如何把nil 转化成一个 指定类型的变量

	   a int = 16
	   b int64 = int64(a)
	   (int64类型)(值)
		  (*AliOssStore)(nil)
*/

// 接口的强制约束
var (
	_ store.Uploader = &AliOssStore{}
)

// new 一个传参的结构体
type Options struct {
	EndPoint     string
	AccessKey    string
	AccessSecret string
}

// 在这结构体上写一个结构体方法做参数校验
func (o *Options) Validate() error {
	if o.EndPoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint,access_key,access_secret,bucket_name has one empty")
	}
	return nil
}

// 默认构造函数, 初始化了Options
func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		EndPoint:     setting.Conf.StorageInfo.OssEndpoint,
		AccessKey:    setting.Conf.StorageInfo.AccessKey,
		AccessSecret: setting.Conf.StorageInfo.AccessSecret,
	})
}

// AliOssStore对象的构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// 检验输入参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	// new 一个客户端
	c, err := oss.New(opts.EndPoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:   c,
		listener: NewDefaultProgressListener(),
	}, nil
}

// 面向对象第一步是构造对象,所有我们先构造一个对象
type AliOssStore struct {
	// 阿里云 OSS client, 私有变量, 不运行外部
	client *oss.Client
	// 依赖listener的实现
	listener oss.ProgressListener
}

// 然后实现我们抽象的接口
func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {

	// 1、获取client, 构造函数中已经实现

	// 2、获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 3、上传文件到bucket
	// objectKey, fileName -----> oss上的文件名, 要上传的文件名
	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
		return err
	}

	// 4、打印下载链接
	download, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL: %s\n", download)
	// fmt.Println("请在一天内下载文件")
	return nil
}
