package server

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type Session struct {
	srv     *Server
	scanner *bufio.Scanner

	conn net.Conn

	clientId int
}

func NewSession(srv *Server, conn net.Conn) *Session {
	return &Session{
		srv:     srv,
		scanner: bufio.NewScanner(conn),
		conn:    conn,
	}
}

func (s *Session) Serve() {
	defer s.close()

	log.Infof("starting session with client: %v", s.conn.RemoteAddr())

	if err := s.auth(); err != nil {
		log.Errorf("session serve: err in auth: %v", err)
		return
	}
	log.Infof("client id: %v success connected", s.clientId)

	for s.scanner.Scan() {
		err := s.srv.Repo.WriteLine(s.clientId, s.scanner.Text())
		if err != nil {
			log.Errorf("can't write to repo err: %v", err)
			return
		}
	}

	if err := s.scanner.Err(); err != nil {
		log.Errorf("session serve: scanner err: %v")
	}

	log.Infof("closing connection with client id: %v", s.clientId)
}

func (s *Session) auth() error {
	if !s.scanner.Scan() {
		return fmt.Errorf("scanner err: %v", s.scanner.Err())
	}

	id, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		return fmt.Errorf("convert to int err: %v", err)
	}

	s.clientId = id

	return nil
}

func (s *Session) Reject() {
	log.Warnf("session reject: too many connections")
	s.close()
}

func (s *Session) close() {
	err := s.conn.Close()
	if err != nil {
		log.Errorf("session connect close err: %v", err)
	}
}
