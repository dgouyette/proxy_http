package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}

func main() {
	port := flag.String("Port", os.Getenv("Port"), "The port which the service will listen on.")
	remoteString := flag.String("Remote", os.Getenv("Remote"), "the target url http://website:8080")

	flag.Parse()

	target, err := url.Parse(*remoteString)

	if err != nil {
		flag.PrintDefaults()
		panic(err)
	}

	log.Printf("forwarding from local port %s to ->  %s %s\n", *port, target.Scheme, target.Host)

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("request = %v", req)
		proxy.ServeHTTP(w, req)
	})

	httpPort := fmt.Sprintf(":%v", *port)

	err = http.ListenAndServe(httpPort, nil)
	if err != nil {
		panic(err)
	}

}
