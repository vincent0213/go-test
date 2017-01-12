package proxy

import (
	"fmt"
	"io"
	"net"
)

const RECV_BUF_LEN = 1024

func Client() {
	listener, err := net.Listen("tcp", "127.0.0.1:9090") //侦听在6666端口
	if err != nil {
		panic("error listening:" + err.Error())
	}
	fmt.Println("Starting the server")
	for {
		conn, err := listener.Accept() //接受连接
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		// n, err := conn.Read(RECV_BUF_LEN)
		// fmt.Println("length:", n)
		fmt.Println("Accepted the Connection :", conn.RemoteAddr(), conn.LocalAddr())
		go EchoServer(conn)
	}

	fmt.Println("###########client##########")

}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Printf("Warning: End of data: %s \n", err)
			return
		default:
			fmt.Printf("Error: Reading data : %s \n", err)
			return
		}
	}
}
