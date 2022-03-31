package session

import (
	"bufio"
	"github.com/amanakin/read-server/pkg/repo"
	"log"
	"net"
	"os"
	"strings"
)

type Session struct {
	file   *os.File
	reader *bufio.Reader

	conn net.Conn

	client uint64
	repo   *repo.Repo
}

func NewSession(file *os.File, conn net.Conn, client uint64, repo *repo.Repo) *Session {
	return &Session{
		file:   file,
		reader: bufio.NewReader(conn),
		conn:   conn,
		client: client,
		repo:   repo,
	}
}

func (session *Session) Serve() {
	defer session.Close()
	for {
		msg, err := session.reader.ReadString('\n')
		msg = strings.TrimRight(msg, "\r\n")
		if err != nil {
			log.Printf("connection read err: %v", err)
			return
		}
		session.repo.WriteLine(session.client, msg)
	}
}

func (session *Session) Reject() {
	log.Printf("too many connections")
	session.Close()
}

func (session *Session) Close() {
	err := session.conn.Close()
	if err != nil {
		log.Printf("session connect close err: %v", err)
	}
}
