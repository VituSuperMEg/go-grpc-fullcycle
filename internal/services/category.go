package services

import (
	"context"

	"github.com/VituSuperMEg/go-grpc-fullcycle/internal/database"
	"github.com/VituSuperMEg/go-grpc-fullcycle/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	categoryID, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

}
