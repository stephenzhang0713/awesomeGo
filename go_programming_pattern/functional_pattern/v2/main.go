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

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

// ServerBuilder 使用一个builder类来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Conf.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxconn int) *ServerBuilder {
	sb.Server.Conf.Maxconns = maxconn
	return sb
}
func (sb *ServerBuilder) WithTimeOut(timeout time.Duration) *ServerBuilder {
	sb.Server.Conf.Timeout = timeout
	return sb
}
func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server.Conf.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() Server {
	return sb.Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server = Server{
		Addr: addr,
		Port: port,
		Conf: &Config{},
	}
	//其它代码设置其它成员的默认值
	return sb
}

func main() {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8080).
		WithProtocol("tcp").
		WithMaxConn(1024).
		WithTimeOut(30 * time.Second).
		Build()

	fmt.Println(server)
}
