package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/crypt"
)

/*
做的加密
*/
func ConfigEncode() (b []byte, err error) {
	key := []byte("学习新思想争做新青年!!")
	i := time.Now().Unix()
	s := strconv.FormatInt(i, 10)
	Data := "@#@113" + s
	fmt.Printf("Data: %v\n", Data)
	b, err = crypt.AesEncrypt([]byte(Data), key)
	return
}
func ConfigDecode(orige []byte) (b []byte, err error) {
	key := []byte("好好学习天天向上冲鸭!!")
	b, err = crypt.AesDecrypt(orige, key)
	return
}
func Send() {
	c, err := net.Dial("tcp", "127.0.0.1:8087")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	c2 := conn.NewConn(c)
	b, err3 := ConfigEncode()
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
		return
	}
	fmt.Printf("b: %v\n", b)
	err2 := c2.WriteLenContent(b) //先写入
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}
func main() {
	c, err := net.Dial("tcp", "127.0.0.1:22123")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	c2 := conn.NewConn(c)
	b, err3 := ConfigEncode()
	/*
		先发送一个加密
	*/
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
		return
	}
	err2 := c2.WriteLenContent(b)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	c2.SetReadDeadlineBySecond(time.Duration(5))
	/*
		应该先发回来一个标志，进行判断是否成功
	*/
	b2, err4 := c2.GetShortLenContent()
	if err4 != nil {
		fmt.Printf("err4: %v\n", err4)
	}
	b3, err5 := ConfigDecode(b2)
	if err5 != nil {
		fmt.Printf("err5: %v\n", err5)
		return
	}
	fmt.Printf("string(b3): %v\n", string(b3))
}
