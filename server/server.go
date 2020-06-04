package server

import (
	"context"

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

// ReadBlog is a unary method to read a blog
func (s *GRPCServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	id := req.GetBlodId()
	blog, err := s.DBHandler.ReadBlog(id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "we can't retrieve blog: %v", err)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       blog.ID.Hex(),
			Title:    blog.Title,
			Content:  blog.Content,
			AuthodId: blog.AuthorID,
		},
	}, nil
}

// ListBlogs is a server streaming function to retrieve all blogs
func (s *GRPCServer) ListBlogs(req *blogpb.ListBlogsRequest, stream blogpb.BlogService_ListBlogsServer) error {
	blogs, err := s.DBHandler.ListBlogs()

	if err != nil {
		return status.Errorf(codes.Internal, "we can't retrieve blogs: %v", err)
	}

	for _, blog := range blogs {
		res := &blogpb.ListBlogsResponse{
			Blog: &blogpb.Blog{
				Id:       blog.ID.Hex(),
				Title:    blog.Title,
				Content:  blog.Content,
				AuthodId: blog.AuthorID,
			},
		}

		if err = stream.Send(res); err != nil {
			return status.Errorf(codes.DataLoss, "someting got wrong to send data: %v", err)
		}
	}

	return nil
}
