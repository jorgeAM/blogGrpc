package main

import (
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/bloGrpc/blogpb"
	"github.com/jorgeAM/bloGrpc/db/mongo"
	"github.com/jorgeAM/bloGrpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	url := os.Getenv("GRPC_SERVER_HOST")
	lis, err := net.Listen("tcp", url)
	defer lis.Close()

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbURL := os.Getenv("DB_URL")
	h, err := mongo.NewDBHandler(dbURL)

	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
	}

	grpcServer := server.GRPCServer{DBHandler: h}
	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &grpcServer)
	reflection.Register(s)
	log.Println("Serving grpc server ...")
	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
