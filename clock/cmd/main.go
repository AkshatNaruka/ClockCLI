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
		Items: []string{"Clock", "Stopwatch", "Alarm", "Time Zone","Calendar","Exit"},
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
		case "Time Zone":
			fmt.Println("Time Zone Utility:")
			fmt.Println("1. Display Current Time")
			fmt.Println("2. Convert Time")
			fmt.Println("Press 'q' to quit")
		
			var option int
			fmt.Scan(&option)
		
			switch option {
			case 1:
				go pkg.DisplayTime()
				fmt.Println("Press 'q' to quit")
				reader := bufio.NewReader(os.Stdin)
				for {
					char, _, err := reader.ReadRune()
					if err != nil {
						fmt.Println(err)
						return
					}
		
					if char == 'q' {
						pkg.StopDisplayTime() // Stop the dynamic display of time when 'q' is pressed
						break
					}
				}
			case 2:
				pkg.ConvertTime()
			default:
				fmt.Println("Invalid option. Please choose 1 or 2.")
			}
		case "Calendar":
			pkg.DisplayCalendar()
		case "Exit":
			fmt.Println("Exiting the CLI...")
			os.Exit(0)
		default:
			fmt.Println("Please choose a utility to run.")
		}
	}
}
			
		
