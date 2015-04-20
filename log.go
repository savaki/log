package log

import (
	"fmt"
	"os"
	"time"

	gklog "github.com/peterbourgon/gokit/log"
)

var DefaultLogger Logger

const (
	LevelKey   = "level"
	MessageKey = "message"
	MethodKey  = "method"
	TimeKey    = "ts"
)

//go:generate stringer -type=LogLevel
type LogLevel int

const (
	Undefined LogLevel = iota
	Trace
	Debug
	Info
	Warn
	Error
)

func init() {
	DefaultLogger = Logger{
		underlying: gklog.NewJSONLogger(os.Stdout),
		level:      Info,
	}
}

type Logger struct {
	underlying gklog.Logger
	level      LogLevel
}

func (l Logger) With(keyvals ...interface{}) Logger {
	return Logger{
		underlying: gklog.With(l.underlying, keyvals...),
		level:      l.level,
	}
}

func (l Logger) WithMethod(name string) Logger {
	return Logger{
		underlying: gklog.With(l.underlying, MethodKey, name),
		level:      l.level,
	}
}

func (l Logger) WithLogLevel(level LogLevel) Logger {
	return Logger{
		underlying: l.underlying,
		level:      level,
	}
}

func (l Logger) Println(format string) {
	l.Infof(format)
}

func (l Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l Logger) Tracef(format string, args ...interface{}) {
	l.printf(Trace, format, args...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.printf(Debug, format, args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.printf(Info, format, args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l.printf(Warn, format, args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.printf(Error, format, args...)
}

func (l Logger) printf(level LogLevel, format string, args ...interface{}) {
	if l.level <= level {
		message := fmt.Sprintf(format, args...)
		l.underlying.Log(TimeKey, time.Now().Format(time.RFC3339), LevelKey, level.String(), MessageKey, message)
	}
}

// -----------------------------------------------------------------------------

func With(keyvals ...interface{}) Logger {
	return DefaultLogger.With(keyvals...)
}

func WithMethod(name string) Logger {
	return DefaultLogger.WithMethod(name)
}

func Printf(format string, args ...interface{}) {
	DefaultLogger.Printf(format, args...)
}

func Println(format string) {
	DefaultLogger.Println(format)
}

func Tracef(format string, args ...interface{}) {
	DefaultLogger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}
