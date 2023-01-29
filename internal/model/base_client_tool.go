package model

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type ChannelMessage struct {
	Head       []byte
	Body       []byte
	BodyLength int
	StatusCode string
}

func write() {
	channelMessage := &ChannelMessage{}
	// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 主动写数据给服务器
	conn.Write(channelMessage.msg2byte())

	// 接收服务器回发的数据
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	//n, err := conn.Read()
	//if err != nil {
	//	fmt.Println("conn.Read err:", err)
	//	return
	//}
	//fmt.Println("服务器回发：", string(buf[:n]))
	buf.Bytes()
}

func (cl *ChannelMessage) msg2byte() []byte {
	head := cl.Head
	headLen := len(head)
	body := cl.Body
	bodyLen := len(body)
	return fmt.Append(int32ToBytes(int32(headLen)), head, int32ToBytes(int32(bodyLen)), body)
}

func (cl *ChannelMessage) byte2msg() {

}

func int32ToBytes(i int32) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}
func bytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}
