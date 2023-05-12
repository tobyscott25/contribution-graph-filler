package dates

import (
	"fmt"
	"time"
)

func AskUserForDate() (time.Time, error) {
	var dateString string
	fmt.Print("What date did you start? (DD-MM-YYYY): ")
	fmt.Scanln(&dateString)

	parsedDate, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	return parsedDate, nil
}

func ParseEpoch(epoch int64) time.Time {
	// Convert the epoch to a time.Time value
	return time.Unix(epoch, 0)
}

func HumanReadableDate(dateTime time.Time) string {
	// Format to be human-readable
	return dateTime.UTC().Format("2 January 2006 15:04:05 MST")
}
