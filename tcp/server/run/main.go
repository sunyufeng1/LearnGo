package main

import (
	"log"

	"github.com/sunyufeng1/LearnGo/tcp/server/objS"
)

// linux peizhi  GOOS = linux  GOARCH = amd64
func main() {
	log.Println("tcp s 启动")
	serviceObj := new(objS.TcpServiceObj)
	serviceObj.Run()
	for {

	}
	log.Println("tcp s 进程结束")
}
