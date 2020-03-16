/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/3/16
   Description :
-------------------------------------------------
*/

package zlog2

// 日志写入之前的中间件, 如果 cancel 则不会再传入下一个中间件且不会写入日志
type BeforeHandler func(level Level, v ...interface{}) (cancel bool)

// 日志写入之后的中间件
type AfterHandler func(level Level, v ...interface{})
