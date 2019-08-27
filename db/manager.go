package db

import (
	"fmt"
	"sync"
)

// Manager 管理
type Manager struct {
	clients sync.Map
	configs sync.Map
	l       logger
}

// NewManager 返回一个manager对象
func NewManager(configs map[string]*Config) *Manager {
	mgr := &Manager{}
	mgr.l = defaultLogger{}
	mgr.Store(configs)
	return mgr
}

// NewManagerWithLogger 返回一个manager对象
func NewManagerWithLogger(configs map[string]*Config, l logger) *Manager {
	mgr := &Manager{}
	mgr.l = l
	mgr.Store(configs)
	return mgr
}

// SetLogger 设置日志
func (mgr *Manager) SetLogger(l logger) {
	mgr.l = l
}

// Store 添加配置
func (mgr *Manager) Store(configs map[string]*Config) {
	if configs == nil {
		return
	}

	for name, config := range configs {
		mgr.Add(name, config)
	}
}

// Load 返回一个db对象
func (mgr *Manager) Load(name string) (client *Client, err error) {
	var value interface{}
	var ok bool
	value, ok = mgr.clients.Load(name)
	if !ok {
		value, ok = mgr.configs.Load(name)
		if !ok {
			return nil, fmt.Errorf("%s configuration does not exist", name)
		}

		config, ok := value.(*Config)
		if ok {
			client, err := NewClientWithLogger(config, mgr.l)
			if err == nil {
				mgr.clients.Store(name, client)
			}

			return client, err
		}

		err = fmt.Errorf("%s db client nil", name)
		return
	}

	client, ok = value.(*Client)
	if !ok {
		err = fmt.Errorf("%s db client nil", name)
		return
	}

	return
}

// Delete 删除一个
func (mgr *Manager) Delete(name string) (err error) {
	var value interface{}
	var ok bool
	// 删除配置
	mgr.configs.Delete(name)
	value, ok = mgr.clients.Load(name)
	if ok {
		client, ok := value.(*Client)
		if ok {
			err = client.Close()
		}
	}

	mgr.clients.Delete(name)
	return
}

// Add 添加一个
func (mgr *Manager) Add(name string, config *Config) {
	config.InitWithDefaults()
	mgr.configs.Store(name, config)
}
