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
    "github.com/kataras/golog"
    "gopkg.in/natefinch/lumberjack.v2"
    "io"
    "log"
    "os"
    "runtime"
    "strings"
)

var DefaultLogger = func() *golog.Logger {
    l := golog.New()
    l.Level = golog.DebugLevel
    l.TimeFormat = "2006-01-02 15:04:05"
    l.Printer.IsTerminal = true
    return l
}

func New(conf LogConfig) *golog.Logger {
    return NewWithLogger(golog.New(), conf)
}

func NewWithLogger(log *golog.Logger, conf LogConfig) *golog.Logger {
    log.Level = parserLogLevel(conf.Level)

    if conf.TimeFormat != "" {
        log.SetTimeFormat(conf.TimeFormat)
    }

    if conf.WriteToFile {
        log.AddOutput(makeWriteSyncers(conf))
    }

    log.Printer.IsTerminal = true

    if conf.ShowFileAndLinenum {
        log.Handle(logHandler(conf))
    }

    if conf.ShowInitInfo {
        log.Info("zlog2 初始化成功")
    }
    return log
}

func logHandler(conf LogConfig) golog.Handler {
    return func(l *golog.Log) bool {
        if !conf.InfoLeveNoLinenum || l.Level != golog.InfoLevel {
            _, filename, line, ok := runtime.Caller(5 + conf.CallerSkip)
            if !ok {
                filename, line = "-", 0
            }
            l.Message = fmt.Sprintf("[%s:%d] %s", filename, line, l.Message)
        }
        return false
    }
}

func parserLogLevel(level string) golog.Level {
    l, ok := map[string]golog.Level{
        "disable": golog.DisableLevel,
        "debug":   golog.DebugLevel,
        "info":    golog.InfoLevel,
        "warn":    golog.WarnLevel,
        "error":   golog.ErrorLevel,
        "fatal":   golog.FatalLevel,
    }[strings.ToLower(level)]
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

    //构建lumberjack的hook
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
