package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// Client 数据库客户端
type Client struct {
	*gorm.DB
	*Config
}

type Logger interface {
	Print(v ...interface{})
}

type defaultLogger struct {
}

var defaultlogger = defaultLogger{}

func (l defaultLogger) Print(v ...interface{}) {
	log.Print(strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

// NewClient 创建一个数据库客户端
func NewClient(config *Config) (client *Client, err error) {
	client, err = NewClientWithLogger(config, defaultlogger)
	return
}

// NewClientWithLogger 创建一个数据库客户端
func NewClientWithLogger(config *Config, l Logger) (client *Client, err error) {
	config.InitWithDefaults()
	db, err := gorm.Open(config.Driver, config.DSN)
	if err != nil {
		return
	}

	if config.LogModeSwitch {
		db.LogMode(true)
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
