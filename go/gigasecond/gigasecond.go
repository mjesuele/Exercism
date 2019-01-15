// Package gigasecond contains a function to add a gigasecond to a given time
package gigasecond

import "time"

// AddGigasecond adds a gigasecond (10^9 second) to a given `time.Time` instance
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
