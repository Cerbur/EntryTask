package transport

import (
	"encoding/binary"
	"io"
	"net"
)

// Transport 负责数据传输encode/decode的结构体
type Transport struct {
	conn net.Conn
	Id   int
}

// NewTransport return 一个绑定 conn 的 *Transport
func NewTransport(conn net.Conn, id int) *Transport {
	return &Transport{conn: conn, Id: id}
}

const headerLength = 4

// Send 发送byte Slice到connect中工具func error 暴露出来
func (t *Transport) Send(data []byte) error {
	// 创建一个 header + data 长度的
	buf := make([]byte, headerLength+len(data))

	// 塞header到buffer中
	binary.BigEndian.PutUint32(buf[:headerLength], uint32(len(data)))
	// 塞data到buff中
	copy(buf[headerLength:], data)
	// 将buff写入到connect中
	if _, err := t.conn.Write(buf); err != nil {
		return err
	}

	return nil
}

// Read 读取connect中的数据到byte Slice中，返回byte 和 error暴露
func (t *Transport) Read() ([]byte, error) {
	// 读取头
	header := make([]byte, headerLength)

	// 读取头获取data长度
	if _, err := io.ReadFull(t.conn, header); err != nil {
		return nil, err
	}

	// 获取data长度
	dataLen := binary.BigEndian.Uint32(header)

	// 读取data数据
	data := make([]byte, dataLen)
	if _, err := io.ReadFull(t.conn, data); err != nil {
		return nil, err
	}

	return data, nil
}
