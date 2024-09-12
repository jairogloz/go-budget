package transaction

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/utils"
	"log"
)

// List retrieves a list of transactions for a user within a specified date range and with pagination options.
// It parses the input parameters, validates them, and then queries the repository for the transactions.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - user: The user for whom the transactions are being retrieved.
//   - from: The start date for the transaction list in ISO 8601 format.
//   - to: The end date for the transaction list in ISO 8601 format.
//   - limit: The maximum number of transactions to retrieve.
//   - offset: The number of transactions to skip for pagination.
//   - listType: The type of list to retrieve (currently only supports an empty string).
//
// Returns:
//   - result: The list of transactions or an error message.
//   - err: An error if any of the input parameters are invalid or if there is an issue retrieving the transactions.
func (s Service) List(ctx context.Context, user *core.User, from, to, limit, offset, listType string) (result interface{}, err error) {

	fromAsTime, err := utils.ParseISO8601(from)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid from: %s", core.ErrorValidation, err.Error())
	}
	toAsTime, err := utils.ParseISO8601(to)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid to: %s", core.ErrorValidation, err.Error())
	}
	limitAsInt, err := core.ParsePositiveInt(limit)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid limit: %s", core.ErrorValidation, err.Error())
	}
	offsetAsInt, err := core.ParsePositiveInt(offset)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid offset: %s", core.ErrorValidation, err.Error())
	}

	switch listType {

	case "":
		// List transactions normally
		txs, err := s.txRepo.List(ctx, user.ID, fromAsTime, toAsTime, limitAsInt, offsetAsInt)
		if err != nil {
			log.Println("error listing transactions:", err)
			return nil, fmt.Errorf("error listing transactions: %s", err.Error())
		}

		return txs, nil
	default:
		return nil, fmt.Errorf("%w: invalid list type: %s", core.ErrorValidation, listType)
	}
}
