package order

import (
	"context"

	"github.com/kvii/pkg/database"
)

type Service struct {
	DB *database.DB
}

func (s *Service) Create(ctx context.Context, name string) (string, error) {
	return "Create " + name, nil
}
