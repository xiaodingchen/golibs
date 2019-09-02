package redis

import (
	"fmt"
	"sync"

	"github.com/xiaodingchen/golibs/logger"
)

// Manager 管理
type Manager struct {
	configs sync.Map
	clients sync.Map
	l       logger.Interface
}

// NewManager 创建一个管理
func NewManager(cfgs map[string]*Config) *Manager {
	mgr := NewManagerWithLogger(cfgs, defaultlogger)
	return mgr
}

// NewManagerWithLogger 返回一个manager对象
func NewManagerWithLogger(configs map[string]*Config, l logger.Interface) *Manager {
	mgr := &Manager{}
	mgr.l = l
	mgr.Store(configs)
	return mgr
}

// SetLogger 设置日志
func (mgr *Manager) SetLogger(l logger.Interface) {
	mgr.l = l
}

// Store 添加多个
func (mgr *Manager) Store(cfgs map[string]*Config) {
	if cfgs == nil || len(cfgs) == 0 {
		return
	}

	for k, v := range cfgs {
		mgr.Add(k, v)
	}
}

// Load 读取一个
func (mgr *Manager) Load(name string) (client *Client, err error) {
	if len(name) == 0 {
		err = fmt.Errorf("redis client name empty")
		return
	}

	var v interface{}
	var ok bool
	v, ok = mgr.clients.Load(name)
	if !ok {
		v, ok = mgr.configs.Load(name)
		if !ok {
			return nil, fmt.Errorf("%s redis configuration does not exist", name)
		}

		config, ok := v.(*Config)
		if ok {
			client = newClient(config)
			if client != nil {
				mgr.clients.Store(name, client)
			}

			return client, nil
		}

		err = fmt.Errorf("%s redis client nil", name)
		return
	}

	client, ok = v.(*Client)
	if !ok {
		err = fmt.Errorf("%s redis client nil", name)
		return
	}

	return
}

// Add 添加一个
func (mgr *Manager) Add(name string, config *Config) (err error) {
	if len(name) == 0 {
		err = fmt.Errorf("redis client name empty")
		return
	}

	if config == nil {
		err = fmt.Errorf("%s config nil", name)
		return
	}

	if len(config.Name) == 0 {
		config.Name = name
	}

	config.InitWithDefaults()
	mgr.configs.Store(name, config)
	return
}

// Delete 删除一个
func (mgr *Manager) Delete(name string) (err error) {
	if len(name) == 0 {
		err = fmt.Errorf("redis client name empty")
		return
	}

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

// Clear 清空所有客户端
func (mgr *Manager) Clear() (err error) {
	mgr.configs.Range(func(key, value interface{}) bool {
		var name string
		var ok bool
		name, ok = key.(string)
		if !ok {
			return false
		}
		err = mgr.Delete(name)
		if err != nil {
			mgr.l.Print(err)
			return false
		}
		return true
	})

	return
}
