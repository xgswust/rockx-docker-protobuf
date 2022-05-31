package main

import (
	"context"
	"log"
	"rockx-docker/net/grpc/client"
	"rockx-docker/net/grpc/server"
	"time"
)

func main() {
	println("hello")
	go server.Run()
	time.Sleep(time.Second * 6)

	count, err := client.NewClient("222.212.90.238").GetContainers(context.Background())
	if err != nil {
		return
	}
	log.Println("containers count is: ", count)
}
