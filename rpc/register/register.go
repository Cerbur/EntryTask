package register

import (
	"entrytask/rpc/client"
	"entrytask/rpc/function"
	"entrytask/rpc/server"
	"reflect"
)

// ClientAutoRegister 客户端的自动注册
// golang的反射不是特别了解，明天查阅相关的资料再来实现自动注册功能
func ClientAutoRegister(cli *client.Client, rpcFunction *function.RPCFunction) {
	valueOf := reflect.ValueOf(rpcFunction).Elem()
	typeOf := valueOf.Type()
	numField := valueOf.NumField()
	for i := 0; i < numField; i++ {
		name := typeOf.Field(i).Name
		fn := valueOf.Field(i)
		cli.CallRPC(name, fn.Addr().Interface())
	}
}

// ServerAutoRegister 服务端的自动注册
func ServerAutoRegister(sr *server.RPCServer, rpcFunction *function.RPCFunction) {
	valueOf := reflect.ValueOf(rpcFunction).Elem()
	typeOf := valueOf.Type()
	numField := valueOf.NumField()
	for i := 0; i < numField; i++ {
		name := typeOf.Field(i).Name
		fn := valueOf.Field(i).Interface()
		sr.Register(name, fn)
	}
}
