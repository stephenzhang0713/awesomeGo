package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"kitex/kitex_gen/api"
	"kitex/kitex_gen/api/echoservice"
	"log"
	"time"
)

func main() {
	c, err := echoservice.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		req := &api.Request{Message: "my request"}
		resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
