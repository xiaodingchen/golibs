package logger

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

// Logger zerolog logger
type Logger struct {
	logger zerolog.Logger
	config *Config
}

// New 返回一个Logger对象
func New(config *Config) *Logger {
	// 初始化配置
	config.InitWithDefault()
	zerolog.CallerSkipFrameCount = 3
	zerolog.TimeFieldFormat = config.zeroTimeFormat

	w := config.NewWriter()
	var logger zerolog.Logger
	logger = zerolog.New(w).With().Timestamp().Logger()

	if w == ioutil.Discard {
		logger.Level(zerolog.Disabled)
	} else {
		logger.Level(config.zerolevel)
	}

	return &Logger{
		logger: logger,
		config: config,
	}
}

// Log 返回一个 zerlog对象
// 无任何配置
func Log() zerolog.Logger {
	return log.Logger
}

// ZeroLogger 返回一个 zerolog logger
func (l *Logger) ZeroLogger() zerolog.Logger {
	return l.logger
}

// Fields 日志扩展字段
func (l *Logger) Fields(fields map[string]interface{}) *Logger {
	l.logger.UpdateContext(func(ctx zerolog.Context) zerolog.Context {

		return ctx.Fields(fields)
	})

	return l
}

// Hook 钩子注册
func (l *Logger) Hook(h zerolog.Hook) *Logger {
	l.logger = l.logger.Hook(h)
	return l
}

// SetCaller 设置Caller关闭或开启
// caller是一个很耗性能的操作，一般生产环境建议关闭
func (l *Logger) SetCaller(o bool) {
	l.config.Caller = o
}
