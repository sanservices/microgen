package handler

import (
	"function/internal/repository"
)

type Handler struct {
	db repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{
		db: repo,
	}
}
