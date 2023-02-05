package initialize

import (
	"errors"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGormMySQL 初始化并且连接MySQL数据库
// @Author: linzijie(https://github.com/linzijie1998)
func InitGormMySQL() *gorm.DB {
	m := global.DOUYIN_CONFIG.MySQL
	if m.DBName == "" {
		global.DOUYIN_LOGGER.Debug("未指定数据库名或数据库名解析失败")
		panic(errors.New("未指定数据库名或数据库名解析失败"))
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DSN(), // DSN信息
		DefaultStringSize:         256,     // string最大长度
		SkipInitializeWithVersion: false,   // 自动配置MySQL版本
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		global.DOUYIN_LOGGER.Debug(fmt.Sprintf("MySQL数据库连接失败 err: %v", err))
		panic(err)
	}
	return db
}

func AutoMigrateTables() {
	if global.DOUYIN_DB == nil {
		return
	}
	global.DOUYIN_DB.AutoMigrate(&model.UserLogin{})
	global.DOUYIN_DB.AutoMigrate(&model.UserInfo{})
	global.DOUYIN_DB.AutoMigrate(&model.Video{})
	global.DOUYIN_DB.AutoMigrate(&model.Comment{})
}
