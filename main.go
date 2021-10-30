package main

import (
	"context"
	"fmt"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

// This example demonstrates how to use Kademlia to have three peers Alice, Charlie and bob discover
// each other in an open, trustless network.
func main() {
	// Let there be Alice, Bob, and Charlie.

	// alice, err := noise.NewNode()
	// if err != nil {
	// 	panic(err)
	// }

	bob, err := noise.NewNode()
	if err != nil {
		panic(err)
	}

	// charlie, err := noise.NewNode()
	// if err != nil {
	// 	panic(err)
	// }

	// Alice, Bob, and Charlie are following an overlay network protocol called Kademlia to discover, interact, and
	// manage each others peer connections.

	//ka, kb, kc := kademlia.New(), kademlia.New(), kademlia.New()
	kb := kademlia.New()

	//alice.Bind(ka.Protocol())
	bob.Bind(kb.Protocol())
	//charlie.Bind(kc.Protocol())

	// if err := alice.Listen(); err != nil {
	// 	panic(err)
	// }

	if err := bob.Listen(); err != nil {
		panic(err)
	}

	// if err := charlie.Listen(); err != nil {
	// 	panic(err)
	// }

	// Have Bob and Charlie learn about Alice. Bob and Charlie do not yet know of each other.
	// nd, err := bob.Ping(context.TODO(), "192.168.1.117:49972")
	// if err != nil {
	// 	panic(err)
	// }

	err = bob.Send(context.TODO(), "192.168.1.117:49972", []byte("test"))
	if err != nil {
		panic(err)
	}

	// if _, err := charlie.Ping(context.TODO(), bob.Addr()); err != nil {
	// 	panic(err)
	// }

	// Using Kademlia, Bob and Charlie will learn of each other. Alice, Bob, and Charlie should
	// learn about each other once they run (*kademlia.Protocol).Discover().

	//fmt.Printf("Alice discovered %d peer(s).\n", len(ka.Discover()))
	// fmt.Printf(nd.ID().Address)
	// kb.Ack(nd.ID())
	fmt.Printf("Bob discovered %d peer(s).\n", len(kb.Discover()))
	//fmt.Printf("Charlie discovered %d peer(s).\n", len(kc.Discover()))

	// Output:
	// Alice discovered 2 peer(s).
	// Bob discovered 2 peer(s).
	// Charlie discovered 2 peer(s).
}
