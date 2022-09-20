package dataserial

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct {
	Name string
	Args []interface{}
	Err  string
}

func Encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (RPCData, error) {
	var data RPCData
	decoder := gob.NewDecoder(bytes.NewBuffer(b))
	if err := decoder.Decode(&data); err != nil {
		return RPCData{}, err
	}
	return data, nil
}
