// Package main implements a server for Echo.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Shinkomori19/intern/pkg/grpc"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 8000, "The server port")

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) EchoMessage(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{ResMessage: "Echoing what server got..\n" + in.GetReqMessage()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
