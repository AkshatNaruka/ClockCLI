package pkg

import (
    "time"
	"fmt"
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

// Start starts the alarm with blinking effect in the terminal
func (a *Alarm) Start() {
    go func() {
        defer func() {
            a.done <- true
        }()

        for {
            select {
            case <-a.quit:
                return
            default:
                if time.Now().After(a.time) {
                    a.blinkAlarm()
                    return
                }
                remaining := a.time.Sub(time.Now())
                hours := int(remaining.Hours())
                minutes := int(remaining.Minutes()) % 60
                seconds := int(remaining.Seconds()) % 60

                // Only update the display when the alarm is not blinking
                fmt.Printf("\rTime remaining %02d:%02d:%02d", hours, minutes, seconds)
                time.Sleep(1 * time.Second)
            }
        }
    }()
}

func (a *Alarm) blinkAlarm() {
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            fmt.Print("\r\033[KAlarm!          ")
            time.Sleep(500 * time.Millisecond)
            fmt.Print("\r\033[K               ")
        case <-a.quit:
            fmt.Println("\r\033[KAlarm!          ") // Print final message before quitting
            return
        case <-a.done:
            fmt.Println("\r\033[KAlarm!          ") // Print final message before returning
            return
        }
    }
}

// Stop stops the alarm
func (a *Alarm) Stop() {
    a.quit <- true
    fmt.Println("Alarm has Stopped")
}