package main

import (
	"entrytask/http/config"
	"entrytask/http/router"
)

func main() {
	_ = config.Fun // RPC函数
	_ = config.Cli // RPC客户端
	router.GetHttpServer().Run()
}
