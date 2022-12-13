// Package main implements a client for Echo.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Shinkomori19/intern/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultReqMessage = "No specified message."
)

var (
	addr       = flag.String("addr", "localhost:50051", "the address to connect to")
	reqMessage = flag.String("reqMessage", defaultReqMessage, "Message sent to server")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.EchoMessage(ctx, &pb.EchoRequest{ReqMessage: *reqMessage})
	if err != nil {
		log.Fatalf("Could not echo : %v", err)
	}
	log.Printf("Greeting: %s", r.GetResMessage())
}
