package services

import (
	"context"
	"github.com/usmaarn/locstique_api/internal/database"
)

type Service struct {
	db  *database.Queries
	ctx context.Context
}

func NewService(ctx context.Context, db *database.Queries) *Service {
	return &Service{db, ctx}
}
