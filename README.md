# 朴实无华的日志模块

---

# 获得
` go get -u github.com/zlyuancn/zlog2 `

# 简单使用

```go
zlog2.Debug("debug")
zlog2.Info("info")
zlog2.Warn("warn")
zlog2.Error("error")
zlog2.Fatal("fatal")
```

# 将日志输出到文件

```go
conf := zlog2.DefaultConfig
conf.WriteToFile = true
l := zlog2.New(conf)
l.Info("123")
```

# 其他选项

```
Level              string       // 日志等级: disable, debug, info, warn, error, fatal
WriteToFile        bool         // 日志是否输出到文件
Name               string       // 日志文件名, 末尾会自动附加 .log 后缀
AppendPid          bool         // 是否在日志文件名后附加进程号
Path               string       // 日志存放路径
FileMaxSize        int          // 每个日志最大尺寸,单位M
FileMaxBackupsNum  int          // 日志文件最多保存多少个备份
FileMaxDurableTime int          // 文件最多保存多长时间,单位天
TimeFormat         string       // 时间输出格式
IsTerminal         bool         // 是否为控制台模式(控制台会打印彩色日志等级)
ShowInitInfo       bool         // 显示初始化信息
ShowFileAndLinenum bool         // 显示文件路径和行号
InfoLeveNoLinenum  bool         // info等级不显示文件路径和行号
CallerSkip         int          // 程序跳转次数
```
