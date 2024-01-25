package pkg

import "time"

// GetMinutes returns the current minute
func GetMinutes() int {
	return time.Now().Minute()
}