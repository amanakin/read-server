package client

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strconv"
)

type Client struct {
	clientScanner *bufio.Scanner
	serverWriter  *bufio.Writer
}

func NewClient(rd io.Reader) *Client {
	return &Client{
		clientScanner: bufio.NewScanner(rd),
	}
}

func (c *Client) RunProxy(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("can't connect to server: %v", err)
	}

	defer func() {
		fmt.Printf("Closing connect\n")
		err := conn.Close()
		if err != nil {
			log.Errorf("connection close err: %v", err)
		}
	}()

	c.serverWriter = bufio.NewWriter(conn)

	id, err := c.auth()
	if err != nil {
		return fmt.Errorf("auth err: %v", err)
	}

	return c.serve(id)
}

func (c *Client) auth() (int, error) {
	fmt.Printf("Enter your id\n")

	if !c.clientScanner.Scan() {
		return 0, fmt.Errorf("scanner err: %v", c.clientScanner.Err())
	}

	id, err := strconv.Atoi(c.clientScanner.Text())
	if err != nil {
		return 0, fmt.Errorf("atoi err: %v", err)
	}

	return id, nil
}

func (c *Client) serve(id int) error {
	_, err := c.serverWriter.Write([]byte(strconv.Itoa(id) + "\r\n"))
	if err != nil {
		return fmt.Errorf("write to server err: %v", err)
	}
	if err := c.serverWriter.Flush(); err != nil {
		return fmt.Errorf("flush err: %v", err)
	}

	fmt.Printf("Enter message:\n")
	for c.clientScanner.Scan() {
		_, err := c.serverWriter.Write([]byte(c.clientScanner.Text() + "\n"))
		if err != nil {
			return fmt.Errorf("write to server err: %v", err)
		}
		if err := c.serverWriter.Flush(); err != nil {
			return fmt.Errorf("flush err: %v", err)
		}
		fmt.Printf("Enter message:\n")
	}

	if err := c.clientScanner.Err(); err != nil {
		return fmt.Errorf("client scanner err: %v", err)
	}

	return nil
}
