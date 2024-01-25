package pkg

import "time"

// GetHours returns the current hour
func GetHours() int {
	return time.Now().Hour()
}