package main

import (
	"context"
	"kitex/kitex_gen/api"
	"kitex/kitex_gen/api/echoservice"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := echoservice.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		req := &api.Request{Message: "one way"}
		err := c.VisitOneway(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
		if err != nil {
			panic(err)
		}
	}
}
