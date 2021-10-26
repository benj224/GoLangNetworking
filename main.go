package main

import (
	"context"
	"fmt"
	"net"

	"github.com/benj224/GoLangNetworking/networking"
	"github.com/perlin-network/noise"
)

func main() {
	ip := networking.GetPublicIp()
	//reciever()
	//sender()
	fmt.Printf(ip)
}

func reciever() {
	reciever, err := noise.NewNode(noise.WithNodeBindHost(net.ParseIP("192.168.1.114")), noise.WithNodeBindPort(49972))
	if err != nil {
		panic(err)
	}

	defer reciever.Close()

	//var wg sync.WaitGroup

	reciever.Handle(func(ctx noise.HandlerContext) error {
		fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))
		//wg.Done()
		return nil
	})

	if err := reciever.Listen(); err != nil {
		panic(err)
	}

}

func sender() {
	sender, err := noise.NewNode(noise.WithNodeBindHost(net.ParseIP("192.168.1.126")), noise.WithNodeBindPort(49972))
	if err != nil {
		panic(err)
	}

	defer sender.Close()

	if err := sender.Send(context.TODO(), "192.168.1.114:49972", []byte("Hi There")); err != nil {
		panic(err)
	}
}
