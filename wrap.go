/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/1/20
   Description :
-------------------------------------------------
*/

package zlog2

import (
    "fmt"

    "github.com/kataras/golog"
)

type logWrap struct {
    log *golog.Logger
}

func (m *logWrap) Debug(v ...interface{}) {
    m.log.Log(golog.DebugLevel, v...)
}
func (m *logWrap) Info(v ...interface{}) {
    m.log.Log(golog.InfoLevel, v...)
}
func (m *logWrap) Warn(v ...interface{}) {
    m.log.Log(golog.WarnLevel, v...)
}
func (m *logWrap) Error(v ...interface{}) {
    m.log.Log(golog.ErrorLevel, v...)
}
func (m *logWrap) Fatal(v ...interface{}) {
    m.log.Log(golog.FatalLevel, v...)
}

func (m *logWrap) Debugf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.log.Log(golog.DebugLevel, msg)
}
func (m *logWrap) Infof(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.log.Log(golog.InfoLevel, msg)
}
func (m *logWrap) Warnf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.log.Log(golog.WarnLevel, msg)
}
func (m *logWrap) Errorf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.log.Log(golog.ErrorLevel, msg)
}
func (m *logWrap) Fatalf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.log.Log(golog.FatalLevel, msg)
}
