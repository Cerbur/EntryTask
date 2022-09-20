package config

import (
	"entrytask/rpc/register"
	"entrytask/rpc/server"
	"entrytask/tcp/service"
)

const addr = "127.0.0.1:7788"

var RPCServer *server.RPCServer

func init() {
	// rpc router注册
	RPCServer = server.NewServer(addr)
	register.ServerAutoRegister(RPCServer, &service.S)
}
