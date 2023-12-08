package proxy

import (
	"log"

	"github.com/armon/go-socks5"
)

func Socks() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 10080
	addr := ":10080"
	log.Printf("socks5 proxy on %s", addr)
	err = server.ListenAndServe("tcp", addr)
	if err != nil {
		panic(err)
	}
}
