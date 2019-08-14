package main

import (
	"context"
	"log"

	"github.com/xiaodingchen/golibs/logger"

	"github.com/spf13/viper"
	"github.com/xiaodingchen/golibs/config"
)

func main() {
	path := "./configs"
	config.Read(path)
	log.Println(viper.GetStringMapString("redis"))
	log.Println(viper.GetString("service_name"))
	ctx, _ := context.WithCancel(context.Background())
	go config.Watch(path)
	// 日志示例
	logDemo()
	<-ctx.Done()
}

func logDemo() {
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

	l := logger.New(config)

	l.Fields(map[string]interface{}{
		"log": "demo",
	}).Info("logdemo")
}
