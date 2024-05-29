package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	const giga = 1e9

	new_time := t.Add(time.Second * giga)

	return new_time
}
