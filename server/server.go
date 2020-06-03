package main

import (
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/bloGrpc/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	url := os.Getenv("GRPC_SERVER_HOST")
	lis, err := net.Listen("tcp", url)
	defer lis.Close()

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})
	log.Println("Serving grpc server ...")
	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
