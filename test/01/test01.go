package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"ehang.io/nps/lib/conn"
)

func GetLen(s io.Reader) (int, error) {
	var l int32
	err := binary.Read(s, binary.LittleEndian, &l)
	return int(l), err
}
func Handler(c net.Conn) {
	// var a = make([]byte, 1024)
	// var l uint32
	c2 := conn.NewConn(c)
	i, err := c2.GetLen()
	fmt.Printf("err: %v\n", err)
	fmt.Printf("i: %v\n", i)
}
func main() {
	t, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	t2, err2 := net.ListenTCP("tcp", t)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	defer t2.Close()
	for {
		c, err3 := t2.Accept()
		if err3 != nil {
			fmt.Printf("err3: %v\n", err3)
			return
		}
		go Handler(c)
	}
}
