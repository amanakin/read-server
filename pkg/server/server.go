package server

import (
	"github.com/amanakin/read-server/pkg/repo"
	"log"
	"net"
	"strconv"
)

type Server struct {
	maxConnections int
	repo           *repo.Repo
}

func NewServer(repo *repo.Repo, maxConnections int) *Server {
	return &Server{
		maxConnections: maxConn,
		repo:           repo,
	}
}

func (server *Server) ListenAndServe()

func (server *Server) Start(port uint64) {
	ln, err := net.Listen("tcp", ":"+strconv.FormatUint(port, 10))
	if err != nil {
		log.Fatalf("can't start server: %v", err)
	}

	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			log.Printf("listener close error: %v", err)
		}
	}(ln)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("listener accept error: %v", err)
			continue
		}
		go Handler(conn)
	}
}
