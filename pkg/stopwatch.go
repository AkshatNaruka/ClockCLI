package pkg

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Stopwatch struct
type Stopwatch struct {
	start    time.Time
	lastLap  time.Time
	lapCount int
	quit     chan bool
}

// NewStopwatch creates a new Stopwatch
func NewStopwatch() *Stopwatch {
	return &Stopwatch{quit: make(chan bool)}
}

// Start starts the stopwatch
func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.lastLap = s.start // Initialize lastLap with the start time

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println(err)
				return
			}

			if char == 'q' {
				s.quit <- true
				return
			} else if char == '\n' {
				s.recordLap()
			}
		}
	}()

	for {
		select {
		case <-s.quit:
			fmt.Println("\nQuitting the stopwatch...")
			return
		default:
			elapsed := time.Since(s.start)
			hours := int(elapsed.Hours())
			minutes := int(elapsed.Minutes()) % 60
			seconds := int(elapsed.Seconds()) % 60
			fmt.Printf("\r%02d:%02d:%02d", hours, minutes, seconds)
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *Stopwatch) recordLap() {
	s.lapCount++
	elapsedSinceLastLap := time.Since(s.lastLap)
	hours := int(elapsedSinceLastLap.Hours())
	minutes := int(elapsedSinceLastLap.Minutes()) % 60
	seconds := int(elapsedSinceLastLap.Seconds()) % 60
	fmt.Printf("\nLap %d: %02d:%02d:%02d\n", s.lapCount, hours, minutes, seconds)
	s.lastLap = time.Now()
}
