package logger

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
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
	if config.AsyncWriter {
		w = diode.NewWriter(w, config.AsyncSize, time.Duration(config.AsyncInterval)*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
	}

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

// Stack 日志堆栈信息
func (l *Logger) Stack(err error) *Logger {
	l.logger.UpdateContext(func(ctx zerolog.Context) zerolog.Context {
		if err != nil {
			ctx.Err(err)
		}

		return ctx.Stack()
	})

	return l
}
