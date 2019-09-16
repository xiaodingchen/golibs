package redis

import (
	"strings"
	"testing"
)

var fconfig = func() (config *Config) {
	config = &Config{}
	config.Network = "tcp"
	// config.Addr = "127.0.0.1:6379"
	config.Addr = "47.104.187.154:5520"
	config.DB = 0
	return
}

func Test_Client(t *testing.T) {
	config := fconfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatal("redis client err", err)
	}

	if client == nil {
		t.Fatal("redis client nil")
	}

	p, err := client.Ping().Result()
	if err != nil {
		t.Fatal("redis client err", err)
	}

	if strings.ToUpper(p) != Pong {
		t.Fatal("redis client ping nowhere")
	}

	t.Logf("%s redis client pass", client.Name)
}
