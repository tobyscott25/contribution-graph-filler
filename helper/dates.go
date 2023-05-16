package helper

import (
	"fmt"
	"time"
)

func ParseDateInput(dateString string) (time.Time, error) {
	parsedDate, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format")
	}
	return parsedDate, nil
}

func HumanReadableFormat(dateTime time.Time) string {
	return dateTime.UTC().Format("Mon 02 Jan 2006 15:04:05 MST")
}
