package redis

import (
	"fmt"
	"runtime"
	"time"
)

// Config 配置
type Config struct {
	Name         string
	Network      string
	Addr         string
	Password     string
	DB           int
	DialTimeout  time.Duration // 毫秒
	ReadTimeout  time.Duration // 毫秒
	WriteTimeout time.Duration // 毫秒
	PoolSize     int           // 连接池个数
	PoolTimeout  time.Duration // 超时，秒
	MinIdleConns int           // 空闲
	MaxRetries   int           // 重试次数

}

const (
	MaxDialTimeout  = 1000
	MaxReadTimeout  = 1000
	MaxWriteTimeout = 3000
	MaxPoolSize     = 1024
	MaxPoolTimeout  = 2 // second
	MinIdleConns    = 3
	MaxRetries      = 1
)

// InitWithDefaults 初始化redis配置
func (c *Config) InitWithDefaults() {
	cpuNum := runtime.NumCPU()
	if c.DialTimeout <= 0 || c.DialTimeout > time.Duration(cpuNum*MaxDialTimeout) {
		c.DialTimeout = MaxDialTimeout * time.Millisecond
	}

	if c.ReadTimeout <= 0 || c.ReadTimeout > time.Duration(cpuNum*MaxReadTimeout) {
		c.ReadTimeout = MaxReadTimeout * time.Millisecond
	}

	if c.WriteTimeout <= 0 || c.WriteTimeout > time.Duration(cpuNum*MaxWriteTimeout) {
		c.WriteTimeout = MaxWriteTimeout * time.Millisecond
	}

	if c.PoolSize <= 0 || c.PoolSize > MaxPoolSize {
		c.PoolSize = 10 * cpuNum
	}

	if c.PoolTimeout <= 0 || c.PoolTimeout > time.Duration(MaxPoolTimeout*cpuNum) {
		c.PoolTimeout = MaxPoolTimeout * time.Second
	}

	if c.MinIdleConns <= 0 || c.MinIdleConns > MinIdleConns*cpuNum {
		c.MinIdleConns = MinIdleConns
	}

	if c.MaxRetries < 0 || c.MaxRetries > MaxRetries*cpuNum {
		c.MaxRetries = MaxRetries
	}

	if len(c.Name) == 0 {
		c.Name = fmt.Sprintf("%s(%s/%d)", c.Network, c.Addr, c.DB)
	}

}
