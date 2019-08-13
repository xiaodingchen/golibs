package logger

import (
	"fmt"
	"strings"
)

// Debug logs to Debug log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(args ...interface{}) {
	event := l.logger.Debug()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Debugln logs to Debug log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debugln(args ...interface{}) {
	event := l.logger.Debug()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Debugf logs to Debug log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	event := l.logger.Debug()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msgf(format, args...)
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {
	event := l.logger.Info()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Infoln(args ...interface{}) {
	event := l.logger.Info()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	event := l.logger.Info()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msgf(format, args...)
}

// Warn logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warn(args ...interface{}) {
	event := l.logger.Warn()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Warnln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warnln(args ...interface{}) {
	event := l.logger.Warn()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Warnf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, args ...interface{}) {
	event := l.logger.Warn()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msgf(format, args...)
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(args ...interface{}) {
	l.Warn(args...)
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warningln(args ...interface{}) {
	l.Warnln(args...)
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {
	event := l.logger.Error()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Errorln(args ...interface{}) {
	event := l.logger.Error()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	event := l.logger.Error()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msgf(format, args...)
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatal(args ...interface{}) {
	event := l.logger.Fatal()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalln(args ...interface{}) {
	event := l.logger.Fatal()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	event := l.logger.Fatal()
	if l.config.Caller {
		event = event.Caller()
	}
	event.Msgf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *Logger) V(level int) bool {
	return true
}

// Panic logs to Panic log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Panic(args ...interface{}) {
	event := l.logger.Panic()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(fmt.Sprint(args...))
}

// Panicln logs to Panic log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Panicln(args ...interface{}) {
	event := l.logger.Panic()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msg(strings.TrimRight(fmt.Sprintln(args...), "\n"))
}

// Panicf logs to Panic log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Panicf(format string, args ...interface{}) {
	event := l.logger.Panic()
	if l.config.Caller {
		event = event.Caller()
	}

	event.Msgf(format, args...)
}
