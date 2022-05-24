package fileUtils

import (
	"log"
	"net/http"
)

func DownFile() {
	//只有压缩包会被下载 文件等会直接被显示  可下载资源要单独放一个文件夹比较好
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/files/", http.StripPrefix("/files", fs)) //把/files/xx 的请求定位到 ./uploads/xx
	//http.Handle("/uploads/",fs)
	log.Fatal(http.ListenAndServe(":40000", nil))
}
