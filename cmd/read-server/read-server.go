package main

import (
	"flag"
	"github.com/amanakin/read-server/pkg/repo"
	"github.com/amanakin/read-server/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	serverAddr := flag.String("s", "localhost:8080", "address of the read-server")
	filename := flag.String("file", "file.txt", "file where will be written clients data")
	maxConnections := flag.Int("max", 5, "max connections to server")
	isHelp := flag.Bool("help", false, "show help")

	flag.Parse()

	if *isHelp {
		flag.PrintDefaults()
		return
	}

	log.Infof("starting server")

	r, err := repo.NewRepo(*filename)
	if err != nil {
		log.Panicf("can't create repo: %v", err)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Errorf("repo close err: %v", err)
		}
	}()

	srv := server.NewServer(r, *maxConnections)
	srv.RunListener(*serverAddr)
}
