package server

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/bloGrpc/blogpb"
	"github.com/jorgeAM/bloGrpc/db"
	"github.com/jorgeAM/bloGrpc/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCServer implements all methods defined in proto file
type GRPCServer struct {
	DBHandler db.Handler
}

// NewBlog is a unary method to create a new blog
func (s *GRPCServer) NewBlog(ctx context.Context, req *blogpb.NewBlogRequest) (*blogpb.NewBlogResponse, error) {
	blog := req.GetBlog()
	data := models.Blog{
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
		AuthorID: blog.GetAuthodId(),
	}

	b, err := s.DBHandler.NewBlog(data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "we can't create new blog: %v", err)
	}

	return &blogpb.NewBlogResponse{
		Blog: &blogpb.Blog{
			Id:       b.ID.Hex(),
			Title:    b.Title,
			Content:  b.Content,
			AuthodId: b.AuthorID,
		},
	}, nil
}
