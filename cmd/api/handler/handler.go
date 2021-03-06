package handler

import "github.com/cobbinma/example-go-api/pkg/models"

type handler struct {
	repository models.Repository
}

func NewHandler(repo models.Repository) *handler {
	return &handler{repository: repo}
}
