package initialize

import (
	"fmt"

	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/spf13/viper"
)

// InitViper 读取配置文件, 配置仅在调用该方法时读取, 暂不支持写入运行时配置
// @Author: linzijie(https://github.com/linzijie1998)
func InitViper(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error config file: %v\n", err)
		panic(err)
	}
	if err := v.Unmarshal(&global.DOUYIN_CONFIG); err != nil {
		fmt.Printf("Fatal error unmarshal: %v\n", err)
		panic(err)
	}
	return v
}
