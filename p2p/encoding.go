package p2p

import (
	"encoding/gob"
	"io"
)

// 编码将数据转换为可以在网络上传输的格式，而解码则将其转换回原始格式。
type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

// ?
func (dec GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	rpc.Playload = buf[:n]

	return nil
}
