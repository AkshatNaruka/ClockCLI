package main

import (
	"clock/pkg"
	"fmt"
	"github.com/manifoldco/promptui"
	"time"
	"bufio"
	"os"
)


type Pomodoro struct {
	workTime  time.Duration
	breakTime time.Duration
	quit      chan bool
	done      chan bool
}

func NewPomodoro(workTime, breakTime time.Duration) *Pomodoro {
	return &Pomodoro{
		workTime:  workTime,
		breakTime: breakTime,
		quit:      make(chan bool),
		done:      make(chan bool),
	}
}

func (p *Pomodoro) Start() {
	go func() {
		for {
			select {
			case <-p.quit:
				return
			default:
				fmt.Println("Work!")
				time.Sleep(p.workTime)
				err := beeep.Notify("Pomodoro", "Time for a break!", "")
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Break!")
				time.Sleep(p.breakTime)
				err = beeep.Notify("Pomodoro", "Break's over, back to work!", "")
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}()
}

func (p *Pomodoro) Stop() {
	p.quit <- true
	fmt.Println("Pomodoro has stopped")
}