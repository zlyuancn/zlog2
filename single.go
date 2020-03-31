/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/3/31
   Description :
-------------------------------------------------
*/

package zlog2

import (
    "fmt"

    "github.com/kataras/golog"
)

var defaultLog = func() *logWrap {
    l := golog.New()
    l.Level = golog.DebugLevel
    l.TimeFormat = "2006-01-02 15:04:05"
    l.Printer.IsTerminal = true
    l.Handle(logHandler(0, false))
    return &logWrap{log: l}
}()

func Log(level Level, v ...interface{}) {
    defaultLog.print(level, v...)
}
func Debug(v ...interface{}) {
    defaultLog.print(DebugLevel, v...)
}
func Info(v ...interface{}) {
    defaultLog.print(InfoLevel, v...)
}
func Warn(v ...interface{}) {
    defaultLog.print(WarnLevel, v...)
}
func Error(v ...interface{}) {
    defaultLog.print(ErrorLevel, v...)
}
func Fatal(v ...interface{}) {
    defaultLog.print(FatalLevel, v...)
}

func Logf(level Level, format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(level, msg)
}
func Debugf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(DebugLevel, msg)
}
func Infof(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(InfoLevel, msg)
}
func Warnf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(WarnLevel, msg)
}
func Errorf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(ErrorLevel, msg)
}
func Fatalf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    defaultLog.print(FatalLevel, msg)
}

func SetBeforeHandler(befores ...BeforeHandler) {
    defaultLog.SetBeforeHandler(befores...)
}
func AddBeforeHandler(befores ...BeforeHandler) {
    defaultLog.AddBeforeHandler(befores...)
}

func SetAfterHandler(afters ...AfterHandler) {
    defaultLog.SetAfterHandler(afters...)
}
func AddAfterHandler(afters ...AfterHandler) {
    defaultLog.AddAfterHandler(afters...)
}
