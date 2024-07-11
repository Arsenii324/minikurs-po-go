package accounts

import (
	"gRPCProject/accounts/models"
	"sync"
)

func New() *Handler {
	return &Handler{
		Accounts: make(map[string]*models.Account),
		Guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	Accounts map[string]*models.Account
	Guard    *sync.RWMutex
}
