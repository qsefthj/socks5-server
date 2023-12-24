package main

import "github.com/qsefthj/go-socks5-server/proxy"

func main() {
	proxy.Socks()
	go proxy.HttpProxy()
}
