package main

import "github.com/sunyufeng1/LearnGo/testBase"

//type Listener struct{}
//
//func (listener Listener) ServeHTTP(w http.ResponseWriter,req *http.Request) {
//	switch req.URL.Path {
//	case "/down":
//		fmt.Fprintf(w, "fileComm will down")
//	default:
//		fmt.Fprintf(w, "404")
//	}
//}

func main() {
	//service := ":8091"
	//
	//udpAddr, err := net.ResolveUDPAddr("udp4", service)
	//checkError(err)
	//
	//conn, err := net.DialUDP("udp", nil, udpAddr)
	//checkError(err)
	//
	//_, err = conn.Write([]byte("anything"))
	//checkError(err)
	//
	//var buf [512]byte
	//n, err := conn.Read(buf[0:])
	//checkError(err)
	//
	//fmt.Println(string(buf[0:n]))
	//
	//os.Exit(0)

	testBase.TestObjArgs() //对象作为参数 的两种情况
	//testBase.TestExit(0)//把程序结束掉
	//testBase.TestPrint()//打印
	//print(testBase.GetCurrentPath())//运行文件的运行时目录
	//testBase.CheckSameObj()
}

/**
	1  基本语法 o
	2  web o
	3 sockets
    4 http 2.0 o
	5 gprc o
	6 windowUi  no 不如用c#去做
	7 文件读写 o
	8 文件上传 下载 0
	9 aws web
	10 aws server
 	11 aws game
*/
/**
1 生成linux 版本 需要设置参数
linux peizhi  GOOS = linux  GOARCH = amd64

2
// 去掉黑色背景 -ldflags="-H windowsgui"
*/
