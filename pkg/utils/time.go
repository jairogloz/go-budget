package utils

import "time"

// ParseISO8601 parses a date string in ISO 8601 format to a time.Time object.
func ParseISO8601(dateStr string) (*time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}
