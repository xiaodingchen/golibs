package logger

import (
	"strings"
	"testing"

	"github.com/rs/zerolog"

	"github.com/kami-zh/go-capturer"
	"google.golang.org/grpc/grpclog"
)

var funcConfig = func() *Config {
	return &Config{
		OutPut:     "stdout",
		Level:      "debug",
		Caller:     true,
		TimeFormat: "datetime",
	}
}

func Test_Logger(t *testing.T) {
	testLoggerV2(t)
}

func Test_OutPut(t *testing.T) {
	config := funcConfig()
	output := capturer.CaptureOutput(func() {
		logger := New(config)
		logger.Infoln("test_info", "demo")
	})

	lines := strings.Split(output, "\n")
	if strings.Contains(lines[0], "test_info") {
		t.Log("logger info output pass")
		return
	}

	t.Fatal("faild")
}

func Test_Fields(t *testing.T) {
	config := funcConfig()
	output := capturer.CaptureOutput(func() {
		logger := New(config)
		logger.Fields(map[string]interface{}{
			"test_field": "demo",
		}).Debug("logger fields test")
	})

	lines := strings.Split(output, "\n")
	if strings.Contains(lines[0], "test_field") {
		t.Log("logger fields pass")
		return
	}

	t.Fatal("faild")
}

type ServiceHook struct {
}

func (h ServiceHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("service", "test_logger_hook")
}

func Test_Hook(t *testing.T) {
	config := funcConfig()
	output := capturer.CaptureOutput(func() {
		logger := New(config)
		logger.Hook(ServiceHook{})
		logger.Info("test_hook")
	})
	t.Log("stdout", output)
	lines := strings.Split(output, "\n")
	if strings.Contains(lines[0], "test_logger_hook") {
		t.Log("hook pass")
		return
	}

	t.Fatal("faild")
}

func Test_Caller(t *testing.T) {
	config := funcConfig()
	output := capturer.CaptureOutput(func() {
		logger := New(config)
		logger.SetCaller(true)
		logger.Info("test_open_call")
		logger.SetCaller(false)
		logger.Info("test_close_call")
		logger.Info("test_default_call")
		logger.SetCaller(true)
		logger.Info("test_open_call")
	})
	lines := strings.Split(output, "\n")
	t.Log("stdout", lines)

	if strings.Contains(lines[0], zerolog.CallerFieldName) &&
		!strings.Contains(lines[1], zerolog.CallerFieldName) &&
		!strings.Contains(lines[2], zerolog.CallerFieldName) &&
		strings.Contains(lines[3], zerolog.CallerFieldName) {
		t.Log("set caller pass")
		return
	}

	t.Fatal("faild")
}

func testLoggerV2(t *testing.T) {
	config := funcConfig()
	var logger interface{}
	logger = New(config)
	_, ok := logger.(grpclog.LoggerV2)
	if ok {
		t.Log("logv2 pass")
		return
	}

	t.Fatal("faild")
}
