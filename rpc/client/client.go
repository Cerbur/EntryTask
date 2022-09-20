package client

import (
	"entrytask/rpc/dataserial"
	"entrytask/rpc/transport"
	"errors"
	"log"
	"net"
	"reflect"
)

type Client struct {
	addr          string                    // 地址与端口
	transportChan chan *transport.Transport // 用于存放transport
	coreThread    int                       // 核心线程数
	maxThread     int                       // 最大线程数
	countThread   int                       // 记录目前Client中有多少个连接
}

func NewClientWithDispatch(addr string, coreThread int, maxThread int) *Client {
	client := Client{
		addr:          addr,
		coreThread:    coreThread,
		transportChan: make(chan *transport.Transport, maxThread),
	}

	for i := 0; i < client.coreThread; i++ {
		client.newConn(i + 1)
	}

	return &client
}

// newConn 创建一个连接
func (c *Client) newConn(id int) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Fatal("new conn fatal", err)
	}
	newTransport := transport.NewTransport(conn, id)

	c.sendToTransportChan(newTransport)
}

func (c *Client) sendToTransportChan(t *transport.Transport) {
	c.transportChan <- t
}

func (c *Client) getFromTransportChan() *transport.Transport {
	return <-c.transportChan
}

func (c *Client) SendTransport(t *transport.Transport) {
	if t != nil {
		c.sendToTransportChan(t)
	}
}

func (c *Client) GetTransport() *transport.Transport {
	transportChan := c.getFromTransportChan()
	return transportChan
}

func (c *Client) CallRPC(rpcName string, fPtr interface{}) {
	// 根据反射模版为fPtr绑定一个fn
	fn := reflect.ValueOf(fPtr).Elem()

	f := func(req []reflect.Value) []reflect.Value {
		// 从channel中获取一个链接，生成一个Transport
		// conn := c.GetConn()
		// now := time.Now()
		reqTransport := c.GetTransport()

		// 定义一个出现异常的返回函数
		errorHandler := func(err error) []reflect.Value {
			// 生成一个fPtr返回值长度一样的Value Slice
			outArgs := make([]reflect.Value, fn.Type().NumOut())
			// 处理非error的返回值，将他们初始化为0值
			for i := 0; i < len(outArgs)-1; i++ {
				outArgs[i] = reflect.Zero(fn.Type().Out(i))
			}

			// 为error赋值
			outArgs[len(outArgs)-1] = reflect.ValueOf(&err).Elem()

			return outArgs
		}

		// 反射填充入参
		inArgs := make([]interface{}, len(req))
		for i, value := range req {
			inArgs[i] = value.Interface()
		}
		// reflectTime1 := time.Now()

		// 生成rpcData的结构体并序列化
		reqRPC := dataserial.RPCData{
			Name: rpcName,
			Args: inArgs,
		}
		buf, err := dataserial.Encode(reqRPC)
		if err != nil {
			return errorHandler(err)
		}
		// serialTime1 := time.Now()

		// 发送
		// sendTime := time.Now()

		err = reqTransport.Send(buf)
		if err != nil {
			return errorHandler(err)
		}

		// 读取返回值
		respond, err := reqTransport.Read()
		if err != nil {
			return errorHandler(err)
		}
		// id := reqTransport.Id
		c.sendToTransportChan(reqTransport)

		// end := time.Now()

		respondData, err := dataserial.Decode(respond)
		if err != nil {
			return errorHandler(err)
		}

		if respondData.Err != "" {
			return errorHandler(errors.New(respondData.Err))
		}

		if len(respondData.Args) == 0 {
			respondData.Args = make([]interface{}, fn.Type().NumOut())
		}

		numOut := fn.Type().NumOut()
		outArgs := make([]reflect.Value, numOut)
		for i := 0; i < numOut; i++ {
			// 处理非 error部分
			if i != numOut-1 {
				if respondData.Args[i] == nil {
					// 为nil初始化0值
					outArgs[i] = reflect.Zero(fn.Type().Out(i))
				} else {
					// 不为nil反射取值
					outArgs[i] = reflect.ValueOf(respondData.Args[i])
				}
			} else {
				// 为什么这里用反射获取0值，因为前面已经把有错误的部分处理掉了（非常无语
				outArgs[i] = reflect.Zero(fn.Type().Out(i))
			}
		}
		// reflectTime2 := time.Now()
		return outArgs
	}

	fn.Set(reflect.MakeFunc(fn.Type(), f))
}
