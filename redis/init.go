package redis

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

type defaultLogger struct {
}

var (
	// DefaultManager 默认管理
	DefaultManager *Manager
	defaultlogger  defaultLogger
)

func (l defaultLogger) Print(v ...interface{}) {
	log.Print(strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func init() {
	initDefaultLogger()
	initDefaultManager()
}

func initDefaultLogger() {
	defaultlogger = defaultLogger{}
}

func initDefaultManager() {
	DefaultManager = NewManager(nil)
}
