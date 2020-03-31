/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/8/30
   Description :
-------------------------------------------------
*/

package zlog2

import (
    "fmt"
    "io"
    "log"
    "os"
    "runtime"
    "strings"

    "github.com/kataras/golog"
    "gopkg.in/natefinch/lumberjack.v2"
)

var DefaultLogger Logfer = defaultLog

type Loger interface {
    Log(level Level, v ...interface{})
    Debug(v ...interface{})
    Info(v ...interface{})
    Warn(v ...interface{})
    Error(v ...interface{})
    Fatal(v ...interface{})
}

type Logfer interface {
    Loger
    Logf(level Level, format string, args ...interface{})
    Debugf(format string, args ...interface{})
    Infof(format string, args ...interface{})
    Warnf(format string, args ...interface{})
    Errorf(format string, args ...interface{})
    Fatalf(format string, args ...interface{})
}

func New(conf LogConfig) LogferWrap {
    return NewWithLogger(golog.New(), conf)
}

func NewWithLogger(log *golog.Logger, conf LogConfig) LogferWrap {
    log.Level = parserLogLevel(Level(strings.ToLower(conf.Level)))

    if conf.TimeFormat != "" {
        log.SetTimeFormat(conf.TimeFormat)
    }

    if conf.WriteToFile {
        log.AddOutput(makeWriteSyncers(conf))
    }

    log.Printer.IsTerminal = true

    if conf.ShowFileAndLinenum {
        log.Handle(logHandler(conf.CallerSkip, conf.InfoLeveNoLinenum))
    }

    if conf.ShowInitInfo {
        log.Info("zlog2 初始化成功")
    }
    return &logWrap{log: log}
}

func logHandler(callerSkip int, infoLeveNoLinenum bool) golog.Handler {
    return func(l *golog.Log) bool {
        if !infoLeveNoLinenum || l.Level != golog.InfoLevel {
            _, filename, line, ok := runtime.Caller(6 + callerSkip)
            if !ok {
                filename, line = "-", 0
            }
            l.Message = fmt.Sprintf("[%s:%d] %s", filename, line, l.Message)
        }
        return false
    }
}

func parserLogLevel(level Level) golog.Level {
    l, ok := levelMapping[level]
    if ok {
        return l
    }

    return golog.InfoLevel
}

func makeWriteSyncers(conf LogConfig) io.Writer {
    // 创建文件夹
    err := os.MkdirAll(conf.Path, 666)
    if err != nil {
        log.Fatal(fmt.Sprintf("无法创建日志目录: <%s>: %s", conf.Path, err))
    }

    // 构建lumberjack的hook
    name := conf.Name
    if conf.AppendPid {
        name = fmt.Sprintf("%s_%d", name, os.Getpid())
    }
    lumberjackHook := &lumberjack.Logger{
        Filename:   fmt.Sprintf("%s/%s.log", conf.Path, name), // 日志文件路径
        MaxSize:    conf.FileMaxSize,                          // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: conf.FileMaxBackupsNum,                    // 日志文件最多保存多少个备份
        MaxAge:     conf.FileMaxDurableTime,                   // 文件最多保存多少天
        Compress:   false,                                     // 是否压缩
    }
    return lumberjackHook
}
