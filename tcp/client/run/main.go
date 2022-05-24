package main

import (
	"log"

	"github.com/sunyufeng1/LearnGo/tcp/client/objC"
)

func main() {
	log.Println("tcp client begin")
	clientObj := new(objC.ClientObj)
	clientObj.Run()
	for {

	}
	log.Println("tcp client end")
}
