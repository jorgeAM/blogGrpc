package main

import (
	"log"
	"os"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/bloGrpc/blogpb"
)

func main() {
	url := os.Getenv("GRPC_SERVER_HOST")
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	defer cc.Close()

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	blogpb.NewBlogServiceClient(cc)
}
