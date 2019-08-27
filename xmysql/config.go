package xmysql

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/xiaodingchen/golibs/db"
)

// Config MySQL配置
type Config struct {
	Name   string
	Logger db.Logger
	Config *db.Config
}

// InitWithDefaults 初始化
func (config *Config) InitWithDefaults() {
	config.Config.InitWithDefaults()
	dsn, _ := mysql.ParseDSN(config.Config.DSN)
	dsn.WriteTimeout = config.Config.WriteTimeout
	dsn.ReadTimeout = config.Config.ReadTimeout
	dsn.Timeout = config.Config.DialTimeout
	config.Config.DSN = dsn.FormatDSN()
	if len(config.Name) == 0 {
		config.Name = fmt.Sprintf("%s(%s/%s)", config.Config.Driver, dsn.Addr, dsn.DBName)
	}
}
