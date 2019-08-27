package xmysql

import "github.com/xiaodingchen/golibs/db"

// DefaultManager 默认管理
var DefaultManager *db.Manager

func init() {
	initDefaultManager()
}

func initDefaultManager() {
	DefaultManager = db.NewManager(nil)
}
