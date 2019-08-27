package xmysql

import (
	"github.com/xiaodingchen/golibs/db"
)

// NewClient 返回一个客户端
func NewClient(config *Config) (client *db.Client, err error) {
	config.InitWithDefaults()
	DefaultManager.SetLogger(config.Logger)
	err = DefaultManager.Add(config.Name, config.Config)
	if err != nil {
		return
	}
	client, err = DefaultManager.Load(config.Name)
	return
}

// NewMultipleClients 创建多个客户端
func NewMultipleClients(cfgs []*Config) (clients []*db.Client, err error) {
	if len(cfgs) == 0 || cfgs == nil {
		return
	}

	for _, config := range cfgs {
		config.InitWithDefaults()
		client, err := NewClient(config)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return
}

// Select 选择一个客户端
func Select(name string) (client *db.Client, err error) {
	client, err = DefaultManager.Load(name)
	return
}
