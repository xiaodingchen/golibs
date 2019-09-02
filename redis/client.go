package redis

import (
	goredis "github.com/go-redis/redis"
)

// Client 客户端
type Client struct {
	*goredis.Client
	*Config
}

// NewClient 创建一个客户端
func NewClient(config *Config) (client *Client, err error) {
	config.InitWithDefaults()
	err = DefaultManager.Add(config.Name, config)
	if err != nil {
		return
	}

	client, err = DefaultManager.Load(config.Name)
	return
}

func newClient(config *Config) *Client {
	config.InitWithDefaults()
	options := &goredis.Options{
		Network:      config.Network,
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		PoolTimeout:  config.PoolTimeout,
	}

	client := goredis.NewClient(options)

	return &Client{
		Client: client,
		Config: config,
	}
}
