package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/perlin-network/noise"
)

func main() {

	reciever()
	//sender()
	// // Let there be nodes Alice and Bob.

	// alice, err := noise.NewNode()
	// if err != nil {
	// 	panic(err)
	// }

	// bob, err := noise.NewNode()
	// if err != nil {
	// 	panic(err)
	// }

	// // Gracefully release resources for Alice and Bob at the end of the example.

	// defer alice.Close()
	// defer bob.Close()

	// var wg sync.WaitGroup

	// // When Alice gets a message from Bob, print it out.

	// alice.Handle(func(ctx noise.HandlerContext) error {
	// 	fmt.Printf("Got a message from Bob: '%s'\n", string(ctx.Data()))
	// 	wg.Done()
	// 	return nil
	// })

	// // When Bob gets a message from Alice, print it out.

	// bob.Handle(func(ctx noise.HandlerContext) error {
	// 	fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))
	// 	wg.Done()
	// 	return nil
	// })

	// // Have Alice and Bob start listening for new peers.

	// if err := alice.Listen(); err != nil {
	// 	panic(err)
	// }

	// if err := bob.Listen(); err != nil {
	// 	panic(err)
	// }

	// // Have Alice send Bob 'Hi Bob!'

	// if err := alice.Send(context.TODO(), bob.Addr(), []byte("Hi Bob!")); err != nil {
	// 	panic(err)
	// }

	// // Wait until Bob receives the message from Alice.

	// wg.Add(1)
	// wg.Wait()

	// // Have Bob send Alice 'Hi Alice!'

	// if err := bob.Send(context.TODO(), alice.Addr(), []byte("Hi Alice!")); err != nil {
	// 	panic(err)
	// }

	// // Wait until Alice receives the message from Bob.

	// wg.Add(1)
	// wg.Wait()

	// fmt.Println(alice.Addr())

}

func reciever() {
	reciever, err := noise.NewNode(noise.WithNodeBindHost(net.ParseIP("192.168.1.120")))
	if err != nil {
		panic(err)
	}

	defer reciever.Close()

	var wg sync.WaitGroup

	reciever.Handle(func(ctx noise.HandlerContext) error {
		fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))
		wg.Done()
		return nil
	})

	if err := reciever.Listen(); err != nil {
		panic(err)
	}

	// if err := reciever.Send(context.TODO(), "", []byte("Hi There")); err != nil {
	// 	panic(err)
	// }

	fmt.Println(reciever.Addr())

}

func sender() {
	sender, err := noise.NewNode()
	if err != nil {
		panic(err)
	}

	defer sender.Close()

	var wg sync.WaitGroup

	sender.Handle(func(ctx noise.HandlerContext) error {
		fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))
		wg.Done()
		return nil
	})

	if err := sender.Listen(); err != nil {
		panic(err)
	}

	if err := sender.Send(context.TODO(), "192.168.1.120:57582", []byte("Hi There")); err != nil {
		panic(err)
	}
}
