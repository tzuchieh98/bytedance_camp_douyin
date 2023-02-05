package config

import (
	"fmt"
	"strings"
)

type MySQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 数据库地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 数据库端口
	DBName   string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库连接登录名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库连接密码
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // DSN配置信息
}

/**
 * DSN 使用`strings.Builder`拼接DSN(data source name)字符串
 * 关于DSN更多信息请参考: https://github.com/go-sql-driver/mysql#dsn-data-source-name
 * @Author: [linzijie](linzijie1998@126.com)
 */
func (m *MySQL) DSN() string {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	var b strings.Builder
	growSize := len(m.Host) + len(m.Port) + len(m.DBName) + len(m.Username) + len(m.Password) + len(m.Config)
	if len(m.Config) == 0 {
		b.Grow(growSize + 9)
		if _, err := fmt.Fprintf(&b, "%s:%s@tcp(%s:%s)/%s", m.Username, m.Password, m.Host, m.Port, m.DBName); err != nil {
			return ""
		}
	} else {
		b.Grow(growSize + 10)
		if _, err := fmt.Fprintf(&b, "%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DBName, m.Config); err != nil {
			return ""
		}
	}
	return b.String()
}
