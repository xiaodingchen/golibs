package logger

import (
	"errors"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"google.golang.org/grpc/grpclog"
)

var funcConfig = func() *Config {
	return &Config{
		OutPut:     "",
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

func Test_Stack(t *testing.T) {
	config := funcConfig()
	logger := New(config)
	err := errors.New("stack_err")
	logger.Fields(map[string]interface{}{
		"test_field": "demo",
	}).Stack(err).Info("test_stack_err")
	t.Log("logger stack pass")
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
