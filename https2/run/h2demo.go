// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE fileComm.

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

var (
	httpsAddr = ":40000"
	httpAddr  = ":40002"

	hostHTTP  = ":40001"
	hostHTTPS = ":40002"
)

func homeOldHTTP(w http.ResponseWriter, r *http.Request) {
	//reqInfoHandler(w,r)
	fmt.Fprintf(w, `<html>
<body>
<p>http 1.%s</p>
</body></html>`, len(w.Header()))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//reqInfoHandler(w,r)
	fmt.Fprintf(w, `<html>
<body>
<p>http 2.%s</p>
</body></html>`, len(w.Header()))
}

func reqInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	//r.Header.Write(w)
}

func registerHandlers() {

	mux2 := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil: // do not allow HTTP/1.x for anything else
			http.Redirect(w, r, "https://"+hostHTTPS+"/", http.StatusFound)
			return
		}
		if r.ProtoMajor == 1 {
			homeOldHTTP(w, r)
			return
		}
		mux2.ServeHTTP(w, r)
	})
	mux2.HandleFunc("/", home)
}

func serveProd() error {
	log.Printf("running in production mode")
	errc := make(chan error, 2)
	go func() { errc <- http.ListenAndServe(":40003", http.DefaultServeMux) }()
	return <-errc
}

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute

// TODO: put this into the standard library and actually send
// PING frames and GOAWAY, etc: golang.org/issue/14204
func idleTimeoutHook() func(net.Conn, http.ConnState) {
	var mu sync.Mutex
	m := map[net.Conn]*time.Timer{}
	return func(c net.Conn, cs http.ConnState) {
		mu.Lock()
		defer mu.Unlock()
		if t, ok := m[c]; ok {
			delete(m, c)
			t.Stop()
		}
		var d time.Duration
		switch cs {
		case http.StateNew, http.StateIdle:
			d = idleTimeout
		case http.StateActive:
			d = activeTimeout
		default:
			return
		}
		m[c] = time.AfterFunc(d, func() {
			log.Printf("closing idle conn %v after %v", c.RemoteAddr(), d)
			go c.Close()
		})
	}
}

func main() {
	var srv http.Server

	srv.Addr = httpsAddr
	srv.ConnState = idleTimeoutHook()

	registerHandlers()

	//serveProd()

	http2.ConfigureServer(&srv, &http2.Server{})

	//go func() {
	//	println("http server")
	//	log.Fatal(http.ListenAndServe(httpAddr, nil))
	//}()

	go func() {
		println("http2 server")
		log.Fatal(srv.ListenAndServeTLS("ssl/1540475527589.pem", "ssl/1540475527589.key"))
	}()
	select {}
}
