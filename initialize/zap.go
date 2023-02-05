package initialize

import (
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/initialize/internal"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitZapLogger() (logger *zap.Logger) {
	if ok, _ := util.PathExists(global.DOUYIN_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.DOUYIN_CONFIG.Zap.Director)
		_ = os.Mkdir(global.DOUYIN_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.DOUYIN_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
