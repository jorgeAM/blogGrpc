package mongo

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jorgeAM/bloGrpc/db"
	"github.com/jorgeAM/bloGrpc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once      sync.Once
	handler   *dbHandler
	someError error
)

type dbHandler struct {
	client *mongo.Client
}

// NewDBHandler returns new instance of dbHandler
func NewDBHandler(url string) (db.Handler, error) {
	once.Do(func() {
		client, err := mongo.NewClient(options.Client().ApplyURI(url))
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = client.Connect(ctx)
		someError = err
		handler = &dbHandler{
			client: client,
		}
	})

	return handler, someError
}

func (h *dbHandler) NewBlog(blog models.Blog) (*models.Blog, error) {
	c := h.client.Database("blogDb").Collection("blogs")
	res, err := c.InsertOne(context.Background(), blog)

	if err != nil {
		return nil, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, errors.New("We cannot convert to oid")
	}

	blog.ID = oid
	return &blog, nil
}

func (h *dbHandler) ReadBlog(id string) (*models.Blog, error) {
	c := h.client.Database("blogDb").Collection("blogs")
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var blog models.Blog
	err = c.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: oid}}).Decode(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (h *dbHandler) ListBlogs() ([]*models.Blog, error) {
	var blogs []*models.Blog
	c := h.client.Database("blogDb").Collection("blogs")
	ctx := context.Background()
	cur, err := c.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var blog models.Blog
		err = cur.Decode(&blog)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (h *dbHandler) UpdateBlog(blog models.Blog) (*models.Blog, error) {
	// c := h.client.Database("blogDb").Collection("blogs")
	return nil, nil
}

func (h *dbHandler) DeleteBlog(id string) error {
	return nil
}
