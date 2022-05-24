package objS

import (
	"log"
	"net"
)

type TcpServiceObj struct {
	listener net.Listener
}

func (this *TcpServiceObj) Run() {
	listener, err := net.Listen("tcp", "127.0.0.1:8087")
	if err != nil {
		panic(err)
	}
	this.listener = listener
	this.Accept()
}

//Accept 等待客户端连接
func (this *TcpServiceObj) Accept() {
	//关闭接口解除阻塞的Accept操作并返回错误
	defer this.listener.Close()
	//循环等待客户端连接
	log.Printf("Waiting for clients...\n")
	for {
		conn, err := this.listener.Accept() //等待客户端连接
		if err != nil {
			log.Printf("Accept Error: %v\n", err)
		} else {
			remoteAddr := conn.RemoteAddr().String() //获取远程客户端网络地址
			log.Printf("TCP Client %v connected\n", remoteAddr)
		}
		//处理客户端连接
		go this.Handle(conn)
	}
}

//Handle 处理客户端连接
func (this *TcpServiceObj) Handle(conn net.Conn) {
	//获取客户端远程地址
	remoteAddr := conn.RemoteAddr().String()
	//延迟关闭客户端连接
	defer conn.Close()
	//循环接收客户端发送的数据
	buf := make([]byte, 1024)
	for {
		//创建字节切片

		//读取时若无消息协程会发生阻塞
		log.Printf("TCP Client %v read block\n", remoteAddr)
		//读取客户端发送的数据
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("TCP Server Read Error: %v\n", err)
			return //退出协程
		}
		//显示客户端发送的数据到服务器终端
		str := string(buf[:n])
		log.Printf("TCP Client %v send message: %v", remoteAddr, str)
		conn.Write([]byte("dddd"))

	}
}
