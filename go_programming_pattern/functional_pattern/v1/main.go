package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) *Server {
	return &Server{Addr: addr, Port: port, Conf: conf}
}

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

func main() {
	srv1 := NewServer("localhost", 9000, nil)

	fmt.Println(srv1)
	conf := Config{
		Protocol: "TCP",
		Timeout:  60 * time.Second,
	}
	srv2 := NewServer("localhost", 9000, &conf)
	fmt.Println(srv2)
}
