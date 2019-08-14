package logger

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
	"github.com/syyongx/php2go"
)

const (
	// DefaultLogFileName 默认日志名称
	DefaultLogFileName = "app.log"
	// DefaultLogDirName 默认日志目录
	DefaultLogDirName = "logs"
	// DefaultLogAsyncTime 异步时间
	DefaultLogAsyncTime = 0
	// DefaultLogAsyncSize 异步输出长度
	DefaultLogAsyncSize = 1 << 10 // 1024
)

// Config 日志配置
type Config struct {
	OutPut         string // 输出
	Level          string // 日志级别
	Caller         bool   // 是否显示行数
	FileName       string // 日志输出的文件路径，建议是绝对路径
	TimeFormat     string // 日志时间戳格式
	AsyncWriter    bool   // 是否异步写入
	AsyncInterval  int    // 定时，单位毫秒
	AsyncSize      int    // 长度
	zerolevel      zerolog.Level
	zeroTimeFormat string
	zerofilename   string
}

// InitWithDefault 配置
func (c *Config) InitWithDefault() {
	if c == nil {
		return
	}

	// 处理日志等级
	level := strings.ToUpper(c.Level)
	switch level {
	case "DEBUG":
		c.zerolevel = zerolog.DebugLevel
	case "INFO":
		c.zerolevel = zerolog.InfoLevel
	case "WARN", "WARNING":
		c.zerolevel = zerolog.WarnLevel
	case "ERROR":
		c.zerolevel = zerolog.ErrorLevel
	case "FATAL":
		c.zerolevel = zerolog.FatalLevel
	case "PANIC":
		c.zerolevel = zerolog.PanicLevel
	default:
		c.zerolevel = zerolog.Disabled
	}

	// 处理日志时间格式
	timeFormat := strings.ToUpper(c.TimeFormat)

	switch timeFormat {
	case "UNIX":
		c.zeroTimeFormat = zerolog.TimeFormatUnix
	case "UNIXMS":
		c.zeroTimeFormat = zerolog.TimeFormatUnixMs
	case "RFC1123":
		c.zeroTimeFormat = time.RFC1123
	case "RFC1123Z":
		c.zeroTimeFormat = time.RFC1123Z
	case "RFC3339NANO":
		c.zeroTimeFormat = time.RFC3339Nano
	case "RFC3339":
		c.zeroTimeFormat = time.RFC3339
	case "DATETIME":
		c.zeroTimeFormat = "2006-01-02 15:04:05"
	case "SHORTTIME":
		c.zeroTimeFormat = "15:04:05"
	case "DATE":
		c.zeroTimeFormat = "2006-01-02"
	case "SHORTDATE":
		c.zeroTimeFormat = "01-02"
	default:
		c.zeroTimeFormat = "2006-01-02 15:04:05.000"
	}

	// 处理输出文件
	if len(c.FileName) < 1 {
		// 获取当前程序运行的目录
		root, err := os.Getwd()
		if err == nil {
			// 获取日志名称
			c.FileName = filepath.Join(root, DefaultLogDirName, DefaultLogFileName)
		}
	}

	suffix := ".log"
	if !strings.HasSuffix(c.FileName, suffix) {
		c.FileName += suffix
	}

	c.zerofilename = c.FileName

	// 处理异步
	if c.AsyncWriter && c.AsyncInterval > 0 && c.AsyncSize == 0 {
		c.AsyncSize = DefaultLogAsyncSize
	}
}

// NewWriter 起一个输出
func (c *Config) NewWriter() io.Writer {
	var w = ioutil.Discard
	var err error

	switch strings.ToUpper(c.OutPut) {
	case "STDOUT":
		w = os.Stdout

	case "STDERR":
		w = os.Stderr

	case "NIL", "NULL":
		w = ioutil.Discard
	default:
		dirFilename := filepath.Dir(c.zerofilename)
		if !php2go.FileExists(dirFilename) {
			err := os.Mkdir(dirFilename, 0766)
			if err != nil {
				log.Panic().Msgf("create logfile dir %s err: %v", dirFilename, err)
			}
		}

		w, err = os.OpenFile(c.zerofilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Panic().Msgf("open logfile %s err: %v", c.zerofilename, err)
		}
	}

	if c.AsyncWriter {
		w = diode.NewWriter(w, c.AsyncSize, time.Duration(c.AsyncInterval)*time.Millisecond, func(missed int) {
			//fmt.Printf("Logger Dropped %d messages", missed)
			// log.Info().Msgf("Logger Dropped %d messages", missed)
		})
	}

	return w
}
