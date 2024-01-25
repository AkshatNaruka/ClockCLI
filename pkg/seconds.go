package pkg

import "time"

// GetSeconds returns the current second
func GetSeconds() int {
	return time.Now().Second()
}