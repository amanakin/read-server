package main

import (
	"flag"
	"github.com/amanakin/read-server/pkg/client"
	log "github.com/sirupsen/logrus"
	"os"
)

var serverAddr = flag.String("s", "localhost:8080", "address of the read-server")

func main() {
	c := client.NewClient(os.Stdin)
	err := c.RunProxy(*serverAddr)
	if err != nil {
		log.Errorf("client err: %v", err)
	}
}
