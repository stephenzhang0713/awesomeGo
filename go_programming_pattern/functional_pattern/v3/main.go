package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) *Server {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv
}

func main() {
	s1 := NewServer("localhost", 1024)
	s2 := NewServer("localhost", 1024, Protocol("udp"))
	s3 := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
	s4 := NewServer("localhost", 1024, TLS(&tls.Config{}))

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)
}
