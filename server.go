package main

import (
	"net"

	. "github.com/immesys/nb"
)

func doserver() {
	NB("rpcbench", "t", "started")
	l, err := net.Listen("tcp", 4000)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go doconn(conn)
	}
}

func doconn(c net.Conn) {

}
