package pkg

import (
    "fmt"
    "time"
    "os"
    "bufio"
)

// Clock struct
type Clock struct {
    quit chan bool
}

// NewClock creates a new Clock
func NewClock() *Clock {
    return &Clock{quit: make(chan bool)}
}

// Run starts the clock
func (c *Clock) Run() {
    go func() {
        reader := bufio.NewReader(os.Stdin)
        for {
            char, _, err := reader.ReadRune()
            if err != nil {
                fmt.Println(err)
                return
            }

            if char == 'q' {
                c.quit <- true
                return
            }
        }
    }()

    for {
        select {
        case <-c.quit:
            fmt.Println("\nQuitting the clock...")
            return
        default:
            fmt.Printf("\r%02d:%02d:%02d", GetHours(), GetMinutes(), GetSeconds())
            time.Sleep(1 * time.Second)
        }
    }
}