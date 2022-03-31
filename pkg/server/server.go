package server

import (
	"github.com/amanakin/read-server/pkg/repo"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

type Server struct {
	ln             net.Listener
	maxConnections int
	Repo           *repo.Repo
}

func NewServer(repo *repo.Repo, maxConnections int) *Server {
	return &Server{
		maxConnections: maxConnections,
		Repo:           repo,
	}
}

func (srv *Server) RunListener(addr string) {
	var limiter chan struct{}
	if srv.maxConnections > 0 {
		limiter = make(chan struct{}, srv.maxConnections)
	} else {
		limiter = nil
	}

	done := make(chan bool)
	srv.listenAndServe(addr, limiter, done)
	<-done
}

func (srv *Server) listenAndServe(addr string, limiter chan struct{}, done chan bool) {
	defer func() { done <- true }()

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("listen server err: %v addr: %v", err, addr)
	}

	srv.serve(ln, limiter)
}

func (srv *Server) serve(ln net.Listener, limiter chan struct{}) {
	defer srv.close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Warnf("server serve: accept temporary err: %v", err)
				time.Sleep(time.Second)
				continue
			}
			// Non-temporary error
			log.Errorf("server serve: accept err: %v", err)
			return
		}

		sess := NewSession(srv, conn)

		if limiter != nil {
			go func() {
				select {
				case limiter <- struct{}{}:
					sess.Serve()
					<-limiter
				default:
					sess.Reject()
				}
			}()
		} else {
			go func() {
				sess.Serve()
			}()
		}
	}
}

func (srv *Server) close() {
	err := srv.ln.Close()
	if err != nil {
		log.Errorf("server close: listener close err: %v", err)
	}
}
