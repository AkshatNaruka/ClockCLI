package pkg

import (
	"bufio"
	"fmt"
	"os"
	"time"
)


func DisplayCalendar() {
	fmt.Println("Calendar:")
	currentTime := time.Now()
	year, month, day := currentTime.Year(), currentTime.Month(), currentTime.Day()

	fmt.Printf("%s %d, %d\n", month, day, year)
	fmt.Println("Sun Mon Tue Wed Thu Fri Sat")

	// Get the first day of the month and the number of days in the month
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
	daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Print spaces for the days before the first day of the month
	for i := 0; i < int(firstDay); i++ {
		fmt.Print("    ")
	}

	// Print the days of the month
	for i := 1; i <= daysInMonth; i++ {
		fmt.Printf("%3d ", i)

		// Move to the next line after Saturday
		if (int(firstDay)+i)%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("\nPress 'q' to quit")
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		}

		if char == 'q' {
			break
		}
	}
}