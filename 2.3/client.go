package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "lock/lock"
)

const (
	address     = "localhost:8082"
)


func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接: %v", err)
	}
	defer conn.Close()
	
	c := pb.NewHelloClient(conn)
	
	empty :=""
	
	r, err := c.Lock(context.Background(), &pb.Request{Message:empty})
			
	if err != nil {
		log.Fatalf("无法握手: %v", err)
	}

	log.Printf("%s", r )

}
