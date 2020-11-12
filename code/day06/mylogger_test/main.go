package main

import (
	"github.com/zhangliying0113/go/code/day06/mylogger"
)

// 测试我们自己写的日志库
var log mylogger.Logger

func main() {
	log = mylogger.NewConsoleLogger("info")                            // 终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "xiaobai.log", 10*1024) // 文件日志实例
	for {
		log.Debug("这是一条 debug 日志")
		log.Trace("这是一条 trace 日志")
		log.Info("这是一条 info 日志")
		log.Warning("这是一条 warning 日志")
		id := 1000
		name := "小白"
		log.Error("这是一条 error 日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条 fatal 日志")
		//time.Sleep(2 * time.Second)
	}
}
