package datatype

import (
	"strings"
	"time"
)

// Date is a type alias for time.Time.
type Date time.Time

// UnmarshalJSON is a method that implements the json.Unmarshaler interface for the Date type.
// It accepts a byte slice as input and returns an error.
// If the input is "null" or an empty string, it returns nil.
// Otherwise, it splits the input by "T" and takes the first part as the date.
// It then parses the date using the time.Parse function with the format `"`+time.DateOnly+`"`.
// If the parsing is successful, it sets the Date value to the parsed time and returns nil.
// If the parsing fails, it returns the error.
func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		return nil
	}

	dateOnly := string(b)
	dateTime := strings.Split(string(b), "T")
	if len(dateTime) > 1 {
		dateOnly = dateTime[0] + "\""
	}

	parsedTime, err := time.Parse(`"`+time.DateOnly+`"`, dateOnly)
	if err != nil {
		return err
	}

	*d = Date(parsedTime)
	return nil
}
