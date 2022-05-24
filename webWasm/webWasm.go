package webWasm

import "net/http"

func RunGoogleHttpDemo() {
	http.ListenAndServe(":40001", http.FileServer(http.Dir("./asset"))) // 此为案例文件夹目录
}
