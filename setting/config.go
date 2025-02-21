package setting

import (
	"heapdump_watcher/models"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 单例模式
var Conf = new(models.Config)

func InitConf() {
	// 指定要读取的配置文件
	viper.SetConfigType("yaml")             // 指定文件格式
	viper.SetConfigName("heapdump-watcher") // 文件名不用加后缀
	viper.AddConfigPath("conf")             // 路径conf/heapdump-watcher.yaml

	// 判断读取配置文件是否有误
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Read configuration file failed, err:", err)
		return
	}

	// 设置默认值,配置文件不存在的时候才会用这个默认值,存在就不用,默认配置文件一定要在viper.Unmarshal(Conf)之前设置
	// viper.SetDefault("oss.BucketName", "xxx")

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		logrus.Error("unmarshal conf failed, err:", err)
		return
	}
	logrus.Infof("使用的配置文件：%s", viper.ConfigFileUsed())

	// 开启配置文件监控（热重载）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("配置文件 %s 发生修改，事件：%v", e.Name, e.Op)
		// 重新解析配置更新全局变量
		if err := viper.Unmarshal(Conf); err != nil {
			logrus.Errorf("热重载配置失败: %v", err)
			return
		}
		logrus.Info("配置热重载成功")
	})
}
