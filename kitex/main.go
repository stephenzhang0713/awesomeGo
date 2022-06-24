package main

import (
	api "kitex/kitex_gen/api/echoservice"
	"log"
)

func main() {
	svr := api.NewServer(new(EchoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
