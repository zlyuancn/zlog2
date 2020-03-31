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
    "sync"

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
    mx      sync.RWMutex
}

var _ LogferWrap = (*logWrap)(nil)

func (m *logWrap) print(level Level, v ...interface{}) {
    m.mx.RLock()
    for _, before := range m.befores {
        if before(level, v...) {
            return
        }
    }
    m.mx.RUnlock()

    m.log.Log(parserLogLevel(level), v...)

    m.mx.RLock()
    for _, after := range m.afters {
        after(level, v...)
    }
    m.mx.RUnlock()
}
func (m *logWrap) Log(level Level, v ...interface{}) {
    m.print(level, v...)
}
func (m *logWrap) Debug(v ...interface{}) {
    m.print(DebugLevel, v...)
}
func (m *logWrap) Info(v ...interface{}) {
    m.print(InfoLevel, v...)
}
func (m *logWrap) Warn(v ...interface{}) {
    m.print(WarnLevel, v...)
}
func (m *logWrap) Error(v ...interface{}) {
    m.print(ErrorLevel, v...)
}
func (m *logWrap) Fatal(v ...interface{}) {
    m.print(FatalLevel, v...)
}

func (m *logWrap) Logf(level Level, format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(level, msg)
}
func (m *logWrap) Debugf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(DebugLevel, msg)
}
func (m *logWrap) Infof(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(InfoLevel, msg)
}
func (m *logWrap) Warnf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(WarnLevel, msg)
}
func (m *logWrap) Errorf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(ErrorLevel, msg)
}
func (m *logWrap) Fatalf(format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    m.print(FatalLevel, msg)
}

func (m *logWrap) SetBeforeHandler(befores ...BeforeHandler) {
    m.mx.Lock()
    m.befores = append(([]BeforeHandler)(nil), befores...)
    m.mx.Unlock()
}
func (m *logWrap) AddBeforeHandler(befores ...BeforeHandler) {
    m.mx.Lock()
    m.befores = append(m.befores, befores...)
    m.mx.Unlock()
}

func (m *logWrap) SetAfterHandler(afters ...AfterHandler) {
    m.mx.Lock()
    m.afters = append(([]AfterHandler)(nil), afters...)
    m.mx.Unlock()
}
func (m *logWrap) AddAfterHandler(afters ...AfterHandler) {
    m.mx.Lock()
    m.afters = append(m.afters, afters...)
    m.mx.Unlock()
}
