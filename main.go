package main

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/benj224/GoLangNetworking/keys"
)

func main() {
	PrKey, PuKey := keys.GenerateKeyPair(256)
	hashId := sha256.Sum256(keys.PublicKeyToBytes(PuKey))
	node1 := node{
		address:    "192.168.1.117:49972",
		id:         string(hashId[:]),
		publicKey:  PuKey,
		privateKey: PrKey,
	}
	table = append(table, node1)
	Broadcast([]byte("hello world"))
}

var b []byte

type node struct {
	address    string
	id         string
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

var table []node

//adding test nodes

func ExampleListener() {
	l, err := net.Listen("tcp", ":49972")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Message received."))

	fmt.Printf(string(buf))
	conn.Close()
}

func ExampleDialer() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "192.168.1.117:49972")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}
}

func Broadcast(message []byte) {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	for _, node := range table {
		defer cancel()

		conn, err := d.DialContext(ctx, "tcp", node.address)
		if err != nil {
			log.Fatalf("Failed to dial: %v", err)
		}
		defer conn.Close()

		if _, err := conn.Write(message); err != nil {
			log.Fatal(err)
		}
	}
}
