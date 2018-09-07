package main

//client.go

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "zhqGo/gRPCDemo/proto"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) >1 {
		name = os.Args[1]
	}
	//var array = []int32{1,2,3,4,6}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name,Age:25})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s,%v", r.Message,r.AgeMessage)


	myRPCRequest(c)
}
func myRPCRequest(c pb.GreeterClient)  {
	users := []string{"zhq", "sweetsmelon"}
	r, err := c.ZHQTest(context.Background(), &pb.TestRequest{Users: users})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("users is:%v,%s", r.Users[0],r.Users[1])
}
