package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/gorilla/mux"
)

var port int
var verbose bool
var globalPort = 8081

type BProxy struct {
	proxy *goproxy.ProxyHttpServer
}

type ProxyMap map[int]BProxy

var m ProxyMap

func HandleProxyList(res http.ResponseWriter,
	req *http.Request) {

	// return current list of proxies

}

func HandleProxyDelete(res http.ResponseWriter,
	req *http.Request) {

	vars := mux.Vars(req)
	port := vars["port"]
	fmt.Println("delete proxy with port", port)

	// TODO: need shut this proxy down
	// m[port]
}

func HandleProxyCreate(res http.ResponseWriter,
	req *http.Request) {

	// port the listen port
	// proxyUsername
	// proxyPassword
	// bindAddress
	// serverBindAddress
	// useEcc
	// trustAllServers

	queryParams := req.URL.Query()
	queryPort := queryParams["port"]
	log.Println("starting a new proxy")

	if len(queryPort) > 0 {
		log.Println("query port is ", queryPort,
			len(queryPort))
	}

	httpProxy := queryParams["httpProxy"]
	if len(httpProxy) > 0 {
		// need to use this proxy
		// for connect
		log.Println("httpProxy param", httpProxy)

	}

	p := goproxy.NewProxyHttpServer()
	p.Verbose = verbose
	if len(httpProxy) > 0 {
		ch := func(req *http.Request) {
			// do nothing here?
			log.Println("got called in connect handler doing nothing")
		}
		p.NewConnectDialToProxyWithHandler(httpProxy[0], ch)
	}
	proxyPort := globalPort
	globalPort++
	m[proxyPort] = BProxy{proxy: p}
	cs := fmt.Sprintf(":%d", proxyPort)
	log.Println("listening on ", cs)
	go http.ListenAndServe(cs, p)

}

func main() {

	m = make(map[int]BProxy)

	port = 8081 // default port to start picking

	fverbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	if *fverbose {
		verbose = true
	} else {
		verbose = false
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy", HandleProxyList).Methods("GET")
	r.HandleFunc("/proxy", HandleProxyCreate).Methods("POST")
	r.HandleFunc("/proxy/{port}", HandleProxyDelete).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
