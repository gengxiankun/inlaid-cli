package viperloader

import (
	"fmt"
	
	"github.com/spf13/viper"
)

func init() {
	// 加载配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/inlaid")
	viper.SetConfigType("yaml")
	// 查找并读取配置文件
	err := viper.ReadInConfig() 
	// 处理读取配置文件的错误
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
