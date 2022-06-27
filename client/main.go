package main

import (
	"context"
	"flag"
	"log"
	"time"
	"crypto/tls"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "grpc-hello-gateway/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:5443", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("srv.crt", "srv.key")
	if err != nil {
		log.Fatal(err)
	}

	// Set up a connection to the server.
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(
		context.Background(),
		*addr,
		grpc.WithBlock(),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cert)),
	)
	
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

