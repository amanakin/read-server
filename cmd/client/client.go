package main

import (
	"flag"
	"github.com/amanakin/read-server/pkg/client"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	serverAddr := flag.String("s", "localhost:8080", "address of the read-server")
	isHelp := flag.Bool("help", false, "show help")

	flag.Parse()

	if *isHelp {
		flag.PrintDefaults()
		return
	}

	c := client.NewClient(os.Stdin)
	err := c.RunProxy(*serverAddr)
	if err != nil {
		log.Errorf("client err: %v", err)
	}
}
