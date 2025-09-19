package handler

import (
	"github.com/Akshatkm1204/go-boilerplate/internal/server"
	"github.com/Akshatkm1204/go-boilerplate/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{}
}
