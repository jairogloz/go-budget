package auth

import "github.com/jairogloz/go-budget/pkg/domain/ports"

type Handler struct {
	AccessControlService ports.AccessControlService
}

func NewHandler(accessControlService ports.AccessControlService) *Handler {
	return &Handler{
		AccessControlService: accessControlService,
	}
}
