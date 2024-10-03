package auth

import "github.com/jairogloz/go-budget/pkg/domain/ports"

type Handler struct {
	AccessControlService ports.AccessControlService
	JWTSecret            []byte
}

func NewHandler(accessControlService ports.AccessControlService, jwtSecret []byte) *Handler {
	h := &Handler{
		AccessControlService: accessControlService,
		JWTSecret:            jwtSecret,
	}
	if h.AccessControlService == nil {
		panic("nil access control service")
	}
	if h.JWTSecret == nil {
		panic("nil JWT")
	}
	return h
}
