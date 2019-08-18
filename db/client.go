package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"github.com/xiaodingchen/golibs/logger"
)

// Client 数据库客户端
type Client struct {
	*gorm.DB
	*Config
}

type defaultLogger struct {
}

func (l defaultLogger) Print(v ...interface{}) {
	log.Print(v...)
}

// New 创建一个数据库客户端
func New(config *Config) (client *Client, err error) {
	client, err = NewWithLogger(config, defaultLogger{})
	return
}

// NewWithLogger 创建一个数据库客户端
func NewWithLogger(config *Config, l logger.Interface) (client *Client, err error) {
	config.InitWithDefaults()
	db, err := gorm.Open(config.Driver, config.DSN)
	if err != nil {
		return
	}

	db.SetLogger(l)

	if config.MaxOpenConns > 0 {
		db.DB().SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxConnLifeTime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(config.MaxConnLifeTime) * time.Second)
	}

	err = db.DB().Ping()
	if err != nil {
		return
	}

	client = &Client{
		DB:     db,
		Config: config,
	}

	return
}
