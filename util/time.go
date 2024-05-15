package util

import (
	"log"
	"strconv"
	"time"
)

func String2UnixTime(timeInput string) (int64, error) {
	// Layout specifying the format of the input string
	// Note: Go uses a specific reference time (Mon Jan 2 15:04:05 MST 2006) to define format layouts
	layout := "2006-01-02T15:04:05"
	log.Println("[app] [Monitor] [time] add time string: " + timeInput)
	// Parse the string into a time.Time struct in local time zone
	parsedTime, err := time.Parse(layout, timeInput)
	if err != nil {
		return 0, err
	}

	// Convert to UTC if not already
	utcTime := parsedTime.UTC()
	unixTime := utcTime.Unix()
	log.Println("[app] [Monitor] [time] unix time: " + strconv.FormatInt(unixTime, 10))
	return unixTime, nil
}
