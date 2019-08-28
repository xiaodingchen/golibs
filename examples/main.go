package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/xiaodingchen/golibs/db"
	"github.com/xiaodingchen/golibs/xmysql"

	"github.com/spf13/viper"
	"github.com/xiaodingchen/golibs/config"
	"github.com/xiaodingchen/golibs/logger"
)

var (
	defalultLogger *logger.Logger
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func init() {
	initConfig()
	initLogger()
	initDB()
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	// 日志示例
	defalultLogger.Info("logger info")
	// db示例
	client, err := xmysql.Select("master")
	if err != nil {
		defalultLogger.ZeroLogger().With().Err(err)
		return
	}
	user := User{}
	client.First(&user)
	userMap := make(map[string]interface{})
	userMap["user"] = user
	defalultLogger.Fields(userMap).Info("user info")
	// logDemo()
	<-ctx.Done()
}

func initConfig() {
	path := "./configs"
	config.Read(path)
	go config.Watch(path)
}

func initLogger() {
	var config *logger.Config
	config = &logger.Config{
		Level:         viper.GetString("log.level"),
		FileName:      viper.GetString("log.logfile"),
		Caller:        viper.GetBool("log.caller"),
		TimeFormat:    viper.GetString("log.timeFormat"),
		AsyncWriter:   viper.GetBool("log.async"),
		AsyncInterval: viper.GetInt("log.asyncInterval"),
		AsyncSize:     viper.GetInt("log.asyncSize"),
	}

	if len(config.FileName) == 0 {
		config.OutPut = "stdout"
	}

	defalultLogger = logger.New(config)
}

func initDB() {
	config := &xmysql.Config{}
	config.Config = &db.Config{}
	dsnFmt := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	config.Name = "master"
	config.Config.LogMode = true
	config.Config.Driver = viper.GetString("db.driver")
	config.Config.DSN = fmt.Sprintf(dsnFmt,
		viper.GetString("db.master.user"),
		viper.GetString("db.master.password"),
		viper.GetString("db.master.host"),
		viper.GetString("db.master.port"),
		viper.GetString("db.master.dbname"),
		viper.GetString("db.master.charset"),
	)

	_, err := xmysql.NewClient(config)
	if err != nil {
		// defalultLogger.Error("mysql client err:", err.Error())
		defalultLogger.ZeroLogger().With().Err(err)
		panic(err)
	}
}
