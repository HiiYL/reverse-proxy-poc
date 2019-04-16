package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	serveReverseProxy("http://127.0.0.1:5000", res, req)
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func main() {
	go func() {
		for {
			PrintMemUsage()
			time.Sleep(time.Second)
		}
	}()
	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe("127.0.0.1:5001", nil); err != nil {
		panic(err)
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
