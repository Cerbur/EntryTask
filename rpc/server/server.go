package server

import (
	"entrytask/rpc/dataserial"
	"entrytask/rpc/transport"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

// RPCServer 用于绑定地址与function
type RPCServer struct {
	addr     string
	function map[string]reflect.Value
}

// NewServer input address 创建 Server
func NewServer(addr string) *RPCServer {
	return &RPCServer{addr: addr, function: make(map[string]reflect.Value)}
}

// Register 注册RPC函数
func (s *RPCServer) Register(funcName string, fn interface{}) {
	if _, ok := s.function[funcName]; !ok {
		s.function[funcName] = reflect.ValueOf(fn)
	}
}

// Execute 用于执行client调用的function
func (s *RPCServer) Execute(request dataserial.RPCData) dataserial.RPCData {
	// 获取注册的function
	f, ok := s.function[request.Name]

	if !ok {
		e := fmt.Sprintf("rpc function:%s not Registered", request.Name)
		log.Println(e)
		return dataserial.RPCData{
			Name: request.Name,
			Args: nil,
			Err:  e,
		}
	}

	// log.Println("function:", request.Name, " is called")

	// 开始解析参数，获取传入值
	inArgs := make([]reflect.Value, len(request.Args))
	for i, a := range request.Args {
		inArgs[i] = reflect.ValueOf(a)
	}

	// 反射调用函数获取返回的Value
	out := f.Call(inArgs)

	// 获取除了 error 之外的返回值
	respondArgs := make([]interface{}, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		respondArgs[i] = out[i].Interface()
	}

	// 判断是否有error
	var errStr string
	if e, ok := out[len(out)-1].Interface().(error); ok {
		errStr = e.Error()
	}
	// 返回
	return dataserial.RPCData{
		Name: request.Name,
		Args: respondArgs,
		Err:  errStr,
	}
}

// Run 运行服务
func (s *RPCServer) Run() {
	// 监听端口
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("listen on %s err %v\n", s.addr, err)
		return
	}

	log.Println("RPCServer listen on ", s.addr)

	count := 1
	// 循环等待socket连接
	for {
		// 获取到连接
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept err: ", err)
			continue
		}
		log.Println("rpc", ":", count, " client connect")
		// 丢给handle处理
		go s.handleConnect(conn, count)
		count++

	}
}

func (s *RPCServer) handleConnect(conn net.Conn, id int) {
	// 创建一个对应到数据交换到到Transport
	connTransport := transport.NewTransport(conn, id)
	for {
		// 读取request
		request, err := connTransport.Read()
		if err != nil {
			if err != io.EOF {
				log.Println("transport read err:", err)
				return
			}
		}

		// 解码request
		requestData, err := dataserial.Decode(request)
		if err != nil {
			log.Println("request decode error:", err)
			return
		}

		// 执行函数
		respondData := s.Execute(requestData)

		// 获取encode得到buffer
		buf, err := dataserial.Encode(respondData)
		if err != nil {
			log.Println("respond encode error", err)
			return
		}

		// 向connect发送数据
		err = connTransport.Send(buf)
		if err != nil {
			log.Println("send error", err)
		}
	}
}
