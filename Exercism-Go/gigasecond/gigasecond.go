package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1 000 000 000 seconds to argumented time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
