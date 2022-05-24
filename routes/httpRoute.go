package routes

import (
	"fmt"
	"net/http"
)

func newLisntener(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/down/":
		//downloadFile(w,req)
	default:
		fmt.Fprintf(w, "40411")
	}
}
