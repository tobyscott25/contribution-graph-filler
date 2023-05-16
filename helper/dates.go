package helper

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

func HumanReadableFormat(dateTime time.Time) string {
	return dateTime.UTC().Format("Mon 02 Jan 2006 15:04:05 MST")
}
