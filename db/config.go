package db

import "time"

const (
	// MaxOpenConn 最大连接数
	MaxOpenConn = 128
	// MaxIdleConn 最大空闲连接
	MaxIdleConn = 16
	// MaxConnLifeTime 链接最大有效时间
	MaxConnLifeTime = 300 // 单位秒
	// MaxDialTimeout 连接超时
	MaxDialTimeout = 1000 // millisecond
	// MaxReadTimeout 读超时
	MaxReadTimeout = 3000 // millisecond
	// MaxWriteTimeout 写超时
	MaxWriteTimeout = 5000 // millisecond
)

// Config 数据库配置
type Config struct {
	Driver          string
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifeTime int
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

// InitWithDefaults 初始化数据库配置
func (c *Config) InitWithDefaults() {
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = MaxOpenConn
	}

	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = MaxIdleConn
	}

	if c.MaxConnLifeTime <= 0 {
		c.MaxConnLifeTime = MaxConnLifeTime
	}

	if c.DialTimeout <= 0 {
		c.DialTimeout = MaxDialTimeout * time.Millisecond
	}

	if c.ReadTimeout <= 0 {
		c.ReadTimeout = MaxReadTimeout * time.Millisecond
	}

	if c.WriteTimeout <= 0 {
		c.WriteTimeout = MaxWriteTimeout * time.Millisecond
	}
}
