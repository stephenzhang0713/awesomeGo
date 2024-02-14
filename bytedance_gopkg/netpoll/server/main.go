package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"
	"time"
)

func main() {
	network, address := "tcp", ":8080"
	listener, _ := netpoll.CreateListener(network, address)
	eventLoop, _ := netpoll.NewEventLoop(handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithOnConnect(connect),
		netpoll.WithReadTimeout(time.Second),
	)
	// start listen loop ...
	eventLoop.Serve(listener)
}

var _ netpoll.OnPrepare = prepare
var _ netpoll.OnConnect = connect
var _ netpoll.OnRequest = handle
var _ netpoll.CloseCallback = close

func prepare(connection netpoll.Connection) context.Context {
	return context.Background()
}

func close(connection netpoll.Connection) error {
	fmt.Printf("[%v] connection closed\n", connection.RemoteAddr())
	return nil
}

func connect(ctx context.Context, connection netpoll.Connection) context.Context {
	fmt.Printf("[%v] connection established\n", connection.RemoteAddr())
	connection.AddCloseCallback(close)
	return ctx
}

func handle(ctx context.Context, connection netpoll.Connection) error {
	reader, writer := connection.Reader(), connection.Writer()
	defer reader.Release()

	msg, _ := reader.ReadString(reader.Len())
	fmt.Printf("[recv msg] %v\n", msg)

	writer.WriteString(msg)
	writer.Flush()
	return nil

}
