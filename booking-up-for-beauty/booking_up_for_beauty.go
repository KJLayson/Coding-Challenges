package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	time, _ := time.Parse(layout, date)

	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	today := time.Now()
	const layout = "January 2, 2006 15:04:05"
	time, _ := time.Parse(layout, date)

	return today.After(time)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	time, _ := time.Parse(layout, date)

	if time.Hour() >= 12 && time.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
// "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"
	time, _ := time.Parse(layout, date)
	fTime := time.Format("Monday, January 2, 2006, at 15:04")

	ret_string := fmt.Sprintf("You have an appointment on %s.", fTime)

	return ret_string
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	const layout = "2006-01-02 15:04:05"
	date := "2024-09-15 00:00:00"
	time, _ := time.Parse(layout, date)

	return time
}
