package objC

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type ClientObj struct {
}

func (this *ClientObj) Run() {
	//拨号远程地址，简历tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8087")
	ClientHandleError(err, "client conn error")

	//预先准备消息缓冲区
	buffer := make([]byte, 1024)

	//准备命令行标准输入
	reader := bufio.NewReader(os.Stdin)

	for {
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)
		n, err := conn.Read(buffer)
		ClientHandleError(err, "client read error")

		serverMsg := string(buffer[:n])
		fmt.Printf("服务端msg", serverMsg)
		if serverMsg == "bye" {
			break
		}

	}
}

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}

}
