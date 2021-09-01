package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"qishuBook/common"
)

func InitViper() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("[error] Fatal error config file:%v", err))
		return
	}
	// 把读取到的配置信息反序列化到Conf中
	err = viper.Unmarshal(Conf)
	common.HandlerError(err, "[error] viper.Unmarshal failed")

	// 实时监控配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件变更后调用的回调函数
		fmt.Println("[info] cfg file changed:", in.Name)
		err := viper.Unmarshal(Conf)
		common.HandlerError(err, "[error] viper.Unmarshal failed")
	})
	return nil
}
