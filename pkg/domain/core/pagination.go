package core

import (
	"fmt"
	"strconv"
)

// ParsePositiveInt converts a string to an integer and ensures it is non-negative.
// If the string cannot be converted to an integer or if the integer is negative,
// an error is returned.
//
// Parameters:
//
//	s - the string to be parsed
//
// Returns:
//
//	int - the parsed integer if valid
//	error - an error if the string cannot be parsed to an integer or if the integer is negative
func ParsePositiveInt(s string) (int, error) {
	parsedInt, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if parsedInt < 0 {
		return 0, fmt.Errorf("must be greater than 0")
	}
	return parsedInt, nil
}
