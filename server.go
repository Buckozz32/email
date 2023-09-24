package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	 "github.com/myuser/myrepo/proto/helloworld"
)

type server struct{
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}


func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	
	s := grpc.NewServer()
helloworldpb.RegisterGreeterServer(s, &server{})
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}



//приложение клиент.го. у меня почему то не создается еще один файл



package main

import (
	"log"
	"os"
"contex"


	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	
	conn, err := grpc.Dial(address, grpc.WithInsecure())
if err != nil {
		log.Fatalf("did not connect: %v", err)}
defer conn.Close()
	c := pb.NewGreeterClient(conn)

name := defaultName
if len(os.Args) > 1 {
name = os.Args[1]}

r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
if err != nil {
log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.Message)
}
