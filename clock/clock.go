package clock

import (
	"fmt"
)

// With the way this challenge is set up, a struct is 100x easier to work with than time.Time
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {

	return Clock{hour: h, minute: m}.Settle()
}

func (c Clock) Add(m int) Clock {
	total_minutes := (c.hour * 60) + c.minute + m
	new_hours := total_minutes / 60
	new_minutes := total_minutes % 60

	c.hour = new_hours
	c.minute = new_minutes
	c = c.Settle()
	return c
}

func (c Clock) Subtract(m int) Clock {
	total_minutes := (c.hour * 60) + c.minute - m
	new_hours := total_minutes / 60
	new_minutes := total_minutes % 60

	c.hour = new_hours
	c.minute = new_minutes
	c = c.Settle()
	return c
}

// Use the total hours and minutes given to the clock to determine what the final time will be
func (c Clock) Settle() Clock {
	m := c.minute
	h := c.hour

	// Calculate how many hours worth of minutes the Clock has
	for m > 59 || m < 0 {
		switch {
		// Negative minutes subtracts from our total hours
		case m < 0:
			m += 60
			h--
		// Positive minutes adds to our total hours
		case m > 59:
			m -= 60
			h++
		}
	}

	// Rolling over/back to the next/prev day until we have the correct hour
	for h > 23 || h < 0 {
		switch {
		// 24:00 is midnight, which is 00:00 on a clock.
		case h == 24:
			h = 0
		// Roll over to the next day
		case h > 24:
			h -= 24
		// Roll back to the prev day
		case h < 0:
			h += 24
		}
	}
	c.minute = m
	c.hour = h

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
