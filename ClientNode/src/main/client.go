package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Arith struct {
	client *rpc.Client
}

type Args struct {
	A, B int
}

func (t *Arith) Multiply(a, b int) int {
	args := &Args{a, b}
	var reply int
	err := t.client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func main() {

	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := net.Dial("tcp", "54.213.238.143:1337")
	if err != nil {
		log.Fatal("Connecting:", err)
	}

	// Create a struct, that mimics all methods provided by interface.
	// It is not compulsory, we are doing it here, just to simulate a traditional method call.
	arith := &Arith{client: rpc.NewClient(conn)}

	fmt.Println(arith.Multiply(5, 6))
}