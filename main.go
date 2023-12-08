package main

import "github.com/qsefthj/go-socks5-server/proxy"

func main() {
	go proxy.HttpProxy()
	proxy.Socks()
}
