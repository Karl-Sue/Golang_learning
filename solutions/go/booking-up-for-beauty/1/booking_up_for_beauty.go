package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	t,_ := time.Parse(layout, date)
    now := time.Now().UTC()
    return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t,_ := time.Parse(layout, date)
    return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
	t,_ := time.Parse(layout, date)
    formattedTime := t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
    return formattedTime
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    layout := "2006-01-02 15:04:05"
    date := "2025-09-15 00:00:00"
	t,_ := time.Parse(layout, date)
	return t
}
