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

type LogferWrap interface {
    Logfer
    SetBeforeHandler(befores ...BeforeHandler)
    AddBeforeHandler(befores ...BeforeHandler)
    SetAfterHandler(afters ...AfterHandler)
    AddAfterHandler(afters ...AfterHandler)
}

type logWrap struct {
    log     *golog.Logger
    befores []BeforeHandler
    afters  []AfterHandler
}

var _ LogferWrap = (*logWrap)(nil)

func (m *logWrap) Log(level Level, v ...interface{}) {
    for _, before := range m.befores {
        if before(level, v...) {
            return
        }
    }

    m.log.Log(parserLogLevel(level), v...)

    for _, after := range m.afters {
        after(level, v...)
    }
}

func (m *logWrap) Debug(v ...interface{}) {
    m.Log(DebugLevel, v...)
}
func (m *logWrap) Info(v ...interface{}) {
    m.Log(InfoLevel, v...)
}
func (m *logWrap) Warn(v ...interface{}) {
    m.Log(WarnLevel, v...)
}
func (m *logWrap) Error(v ...interface{}) {
    m.Log(ErrorLevel, v...)
}
func (m *logWrap) Fatal(v ...interface{}) {
    m.Log(FatalLevel, v...)
}

func (m *logWrap) Debugf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.Log(DebugLevel, msg)
}
func (m *logWrap) Infof(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.Log(InfoLevel, msg)
}
func (m *logWrap) Warnf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.Log(WarnLevel, msg)
}
func (m *logWrap) Errorf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.Log(ErrorLevel, msg)
}
func (m *logWrap) Fatalf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.Log(FatalLevel, msg)
}

func (m *logWrap) SetBeforeHandler(befores ...BeforeHandler) {
    m.befores = append(([]BeforeHandler)(nil), befores...)
}
func (m *logWrap) AddBeforeHandler(befores ...BeforeHandler) {
    m.befores = append(m.befores, befores...)
}

func (m *logWrap) SetAfterHandler(afters ...AfterHandler) {
    m.afters = append(([]AfterHandler)(nil), afters...)
}
func (m *logWrap) AddAfterHandler(afters ...AfterHandler) {
    m.afters = append(m.afters, afters...)
}
