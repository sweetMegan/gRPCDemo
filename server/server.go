package main

// server.go

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "zhqGo/gRPCDemo/proto"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:"123",AgeMessage:in.Age}, nil
}
//方法名必须与Greeter声明的方法名一致
func (s *server)ZHQTest(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {

	return &pb.TestReply{Users:in.Users},nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}