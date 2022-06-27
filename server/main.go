package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworldpb "grpc-hello-gateway/helloworld"

	//ADD for grpc-gateway
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	cert, err := tls.LoadX509KeyPair("srv.crt", "srv.key")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":5443")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer(opts...)
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})
	// Optional reflection service register for use with grpcurl or grpc_cli
	reflection.Register(s)
	// Serve gRPC Server
	log.Printf("Serving gRPC on %v", lis.Addr())
	//log.Fatal(s.Serve(lis))

	// ADD for grpc-gateway
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	connAddr := "localhost:5443" // cannot be 127.0.0.1:5443 because cert is for localhost
	log.Printf("connecting to %v", connAddr)

	conn, err := grpc.DialContext(
		context.Background(),
		connAddr,
		grpc.WithBlock(),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cert)),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	log.Printf("connected to %v", connAddr)

	gwmux := runtime.NewServeMux()
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":5444",
		Handler: gwmux,
	}
	log.Printf("Serving gRPC-Gateway on %v", gwServer.Addr)
	err = gwServer.ListenAndServeTLS("srv.crt", "srv.key")
	if err != nil {
		log.Fatal(err)
	}
}
