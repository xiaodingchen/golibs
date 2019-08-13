package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/xiaodingchen/golibs/config"
)

func TestRead(t *testing.T) {
	path := "../examples/configs"
	config.ReadConfig(path)
	t.Log(viper.GetStringMapString("redis"))
	t.Log(viper.GetString("service_name"))
}
