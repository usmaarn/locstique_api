package handlers

import (
	"context"
	"github.com/usmaarn/locstique_api/internal/database"
	"github.com/usmaarn/locstique_api/internal/services"
)

type Handler struct {
	ctx     context.Context
	service *services.Service
}

func NewHandler(ctx context.Context, db *database.Queries) *Handler {
	service := services.NewService(ctx, db)
	return &Handler{ctx, service}
}
