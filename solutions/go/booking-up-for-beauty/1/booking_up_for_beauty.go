package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("Failed to parse string as Time!")
	}
	log.Default().Println(t)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("Failed to parse string as Time!")
	}
	log.Default().Println(t)

	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("Failed to parse string as Time!")
	}

	hour := t.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("Failed to parse string as Time!")
	}

	weekday := t.Weekday()
	month := t.Month()
	day := t.Day()
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Date(2025, 9, 15, 0, 0, 0, 0, time.UTC)
	layout := "2006-01-02 15:04:05 -0700 MST"
	tStr := t.Format(layout)

	fT, err := time.Parse(layout, tStr)
	if err != nil {
		log.Fatal("Failed to parse anniversary date as datetime!")
	}
	return fT
}
