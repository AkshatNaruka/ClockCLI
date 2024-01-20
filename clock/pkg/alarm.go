package pkg

import (
	"fmt"
	"time"
)

type Alarm struct {
    time time.Time
    quit chan bool
    done chan bool
}

func NewAlarm(t time.Time) *Alarm {
    return &Alarm{
        time: t,
        quit: make(chan bool),
        done: make(chan bool),
    }
}
// Start starts the alarm
func (a *Alarm) Start() {
    go func() {
        for {
            select {
            case <-a.quit:
                return
            default:
                if time.Now().After(a.time) {
                    fmt.Println("\nAlarm!")
					a.done <- true
                    return
                }
                remaining := a.time.Sub(time.Now())
                hours := int(remaining.Hours())
                minutes := int(remaining.Minutes()) % 60
                seconds := int(remaining.Seconds()) % 60
                fmt.Printf("\rTime remaining %02d:%02d:%02d", hours, minutes, seconds)
                time.Sleep(1 * time.Second)
            }
        }
    }()
}

// Stop stops the alarm
func (a *Alarm) Stop() {
	a.quit <- true
	fmt.Println("Alarm has Stopped")
}