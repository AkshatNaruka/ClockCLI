package main

import (
	"clock/pkg"
	"fmt"
	"github.com/manifoldco/promptui"
	"time"
	"bufio"
	"os"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Utility",
		Items: []string{"Clock", "Stopwatch", "Alarm", "Exit"},
	}

	for {
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Clock":
			clock := pkg.NewClock()
			clock.Run()
		case "Stopwatch":
			stopwatch := pkg.NewStopwatch()
			stopwatch.Start()
		case "Alarm":
			
			var minutes int
			fmt.Println("Enter the number of minutes")
			fmt.Scan(&minutes)
			
			alarmTime := time.Now().Add(time.Duration(minutes) * time.Minute)
			alarm := pkg.NewAlarm(alarmTime)
			alarm.Start()
		
			fmt.Println("Press 'q' to quit")
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _, err := reader.ReadRune()
				if err != nil {
					fmt.Println(err)
					return
				}
		
				if char == 'q' {
					alarm.Stop() // Stop the alarm when 'q' is pressed
					break
				}
			}
		case "Exit":
			fmt.Println("Exiting the CLI...")
			os.Exit(0)
		case "Pomodoro":
			pomodoro := pkg.NewPomodoro(25*time.Minute, 5*time.Minute)
			pomodoro.Start()

			fmt.Println("Press 'q' to quit")
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _, err := reader.ReadRune()
				if err != nil {
					fmt.Println(err)
					return
				}

				if char == 'q' {
					pomodoro.Stop()
					return
				}
			}
		default:
			fmt.Println("Please choose a utility to run.")
		}
	}
}
			
		
