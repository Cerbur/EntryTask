package config

import (
	"entrytask/rpc/client"
	"entrytask/rpc/function"
	"entrytask/rpc/register"
	"fmt"
	"os"
)

const (
	addr = "127.0.0.1:7788"
	path = "/picture/"
)

var (
	Fun  function.RPCFunction
	Cli  *client.Client
	Path string
)

func init() {
	// gob 注册
	str, _ := os.Getwd()
	Path = fmt.Sprint(str, path)
	Cli = client.NewClientWithDispatch(addr, 200, 200)
	register.ClientAutoRegister(Cli, &Fun)
}
