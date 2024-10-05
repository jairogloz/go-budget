package core

import "time"

// CustomTime to handle JSON unmarshalling of time.Time
type CustomTime struct {
	time.Time
}

func (ct *CustomTime) ToTime() *time.Time {
	return &ct.Time
}

// UnmarshalJSON method to parse the date string into a time.Time object
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Define the date format
	const layout = "2006-01-02T15:04:05Z07:00"

	// Remove the quotes from the JSON string
	str := string(b)
	str = str[1 : len(str)-1]

	// Parse the string into a time.Time object
	t, err := time.Parse(layout, str)
	if err != nil {
		return err
	}

	// Assign the parsed time to the CustomTime object
	ct.Time = t
	return nil
}
