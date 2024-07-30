package p2p

import (
	"encoding/gob"
	"io"
)

// 编码将数据转换为可以在网络上传输的格式，而解码则将其转换回原始格式。
type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

// ?
func (dec GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Playload = buf[:n]

	return nil
}
