package main

import (
	"github.com/sunyufeng1/LearnGo/https2"
)

func main() {
	//启动一个端口来监听请求
	//对请求进行路由转发
	//转发到具体的处理

	https2.RunGoogleHttps2Demo()

}