package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	cert , err :=tls.LoadX509KeyPair("localhost.crt","localhost.key")
	if err != nil {
		log.Fatalln(err)
	}
	server := &http.Server{
		Addr: ":9000",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_,_ = fmt.Fprintln(writer,"Hello World")
	})
	log.Fatalln(server.ListenAndServeTLS("",""))
}