package main

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"github.com/xiaodingchen/golibs/config"
)

func main() {
	path := "./configs"
	config.ReadConfig(path)
	log.Println(viper.GetStringMapString("redis"))
	log.Println(viper.GetString("service_name"))
	ctx, _ := context.WithCancel(context.Background())
	go config.WatchConfig(path)
	<-ctx.Done()
}
