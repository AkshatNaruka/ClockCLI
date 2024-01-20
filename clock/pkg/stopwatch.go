package pkg

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Stopwatch struct
type Stopwatch struct {
	start time.Time
	quit  chan bool
}

// NewStopwatch creates a new Stopwatch
func NewStopwatch() *Stopwatch {
	return &Stopwatch{quit: make(chan bool)}
}

// Start starts the stopwatch
func (s *Stopwatch) Start() {
	s.start = time.Now()

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