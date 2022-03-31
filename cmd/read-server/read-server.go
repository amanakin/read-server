package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"net"
)

// TODO
// add Zap
// add middleware
// add user repo
// add user repo database

func main() {
	fmt.Printf("starting session\n")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("session can't listen: %v", err)
	}

	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			log.Printf("ln close error: %v", err)
		}
	}(ln)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("listener accept err: %v", err)
			continue
		}
		go Handler(conn)
	}
}
