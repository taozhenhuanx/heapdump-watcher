package models

// 配置文件对应结构体对象
// 对应yaml的mysql  (Mysql--转成小写对应), 如果有不符合的就用标签去指定 `mapstructure:"db"`
type Config struct {
	// 匿名结构体---OSS信息
	StorageInfo struct {
		AccessKey    string `mapstructure:"access_key"`
		AccessSecret string `mapstructure:"access_secret"`
		OssEndpoint  string `mapstructure:"oss_endpoint"`
		BucketName   string `mapstructure:"bucket_name"`
		StorageType  string `mapstructure:"storage_type"`
	} `mapstructure:"storageInfo"`

	// 匿名结构体---发送告警媒介信息
	AlarmMedium struct {
		WeChatKey     string `mapstructure:"wechat_key"`
		WebhookType   string `mapstructure:"webhook_type"`
		DingTalkToken string `mapstructure:"dingtalk_token"`
	} `mapstructure:"alarmMedium"`

	// 匿名结构体---监听的文件路径
	FilePath struct {
		WatchPath string `mapstructure:"watch_path"`
		KubeConf  string `mapstructure:"kube_conf"`
	} `mapstructure:"filePath"`

	// 邮件信息
	AlarmEmail struct {
		User string   `mapstructure:"user"`
		Pass string   `mapstructure:"pass"`
		Host string   `mapstructure:"host"`
		Port int      `mapstructure:"port"`
		To   []string `mapstructure:"to"`
		Cc   []string `mapstructure:"cc"`
		Bcc  []string `mapstructure:"bcc"`
	} `mapstructure:"alarmEmail"`
}
