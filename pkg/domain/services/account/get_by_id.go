package account

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"log"
	"time"
)

// GetByID retrieves an account by its ID for a given user.
// It creates a context with a timeout to ensure the operation does not hang indefinitely.
// If the account is found, it is returned along with a nil error.
// If an error occurs, an empty account and the error are returned.
//
// Parameters:
//
//	ctx - the context for the request, used for cancellation and deadlines
//	userId - the ID of the user who owns the account
//	id - the ID of the account to retrieve
//
// Returns:
//
//	core.Account - the retrieved account
//	error - an error if one occurred, otherwise nil
func (s Service) GetByID(ctx context.Context, userId, id string) (core.Account, error) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	account, err := s.repo.GetByID(ctxWithTimeout, userId, id)
	if err != nil {
		log.Println("error getting account by ID: ", err.Error())
		return core.Account{}, fmt.Errorf("error getting account by ID: %w", err)
	}

	return account, nil
}
