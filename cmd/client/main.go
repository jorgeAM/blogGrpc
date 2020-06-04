package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/bloGrpc/blogpb"
	"google.golang.org/grpc"
)

func main() {
	url := os.Getenv("GRPC_SERVER_HOST")
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	defer cc.Close()

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	c := blogpb.NewBlogServiceClient(cc)
	newBlog(c)
}

func newBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.NewBlogRequest{
		Blog: &blogpb.Blog{
			Title:    "GRPC is awesome",
			Content:  "#teamGRPC",
			AuthodId: "jorguito",
		},
	}
	res, err := c.NewBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("something wrong when call NewBlog method: %v", err)
	}

	fmt.Println(res.GetBlog())
}