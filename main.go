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

func HandleProxyList(res http.ResponseWriter,
	req *http.Request) {

	// return current list of proxies

}

func HandleProxyDelete(res http.ResponseWriter,
	req *http.Request) {

	vars := mux.Vars(req)
	port := vars["port"]
	fmt.Println("delete proxy with port", port)
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

	log.Println("starting a new proxy")

	// create a new proxy pick a new port
	// how do you pick a port
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose
	http.ListenAndServe(":8080", proxy)
	// TODO: need a place to store the proxies

}

func main() {

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
}
