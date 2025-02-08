package models

// 配置文件对应结构体对象
// 对应yaml的mysql  (Mysql--转成小写对应), 如果有不符合的就用标签去指定 `mapstructure:"db"`
type Config struct {
	// 匿名结构体---OSS信息
	OSS struct {
		AccessKey    string `yaml:"AccessKey"`
		AccessSecret int    `yaml:"AccessSecret"`
		OssEndpoint  string `yaml:"OssEndpoint"`
		BucketName   string `yaml:"BucketName"`
	} `mapstructure:"oss"`

	// 匿名结构体---发送告警媒介信息
	AlarmMedium struct {
		WebhookURL  string `mapstructure:"webhook_url"`
		WebhookType string `mapstructure:"webhook_type"`
	} `mapstructure:"alarmMedium"`

	// 匿名结构体---监听的文件路径
	FilePath struct {
		WatchPath string `mapstructure:"watch_path"`
	} `mapstructure:"filePath"`
}
