package main

import (
	"log"
	"net"
	"net/rpc"
)

type AllowedFunctions interface {
    Multiply(args *Args, reply *int) error
}

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {

	//Creating an instance of struct which implement Arith interface
	//allowedFunctions := new(AllowedFunctions)
	arith := new(Arith)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Arith"
	// The wrapper method here ensures that only structs which implement Arith interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	server.Register(arith)

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	server.Accept(l)
}