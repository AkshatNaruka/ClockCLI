package pkg

import (
    "fmt"
    "time"

    "github.com/manifoldco/promptui"
)

func DisplayTime() {
    location := getTimeZone()

    for {
        currentTime := time.Now().In(location)
        fmt.Printf("Current time in %s: %s\r", location, currentTime.Format("2006-01-02 15:04:05"))
        time.Sleep(1 * time.Second) // Update every second

        select {
        case <-time.After(1 * time.Second):
            // Do nothing, just continue
        case <-quitChannel:
            return
        }
    }
}

var quitChannel = make(chan struct{})

// StopDisplayTime stops the dynamic display of time.
func StopDisplayTime() {
    quitChannel <- struct{}{}
}

// ConvertTime converts a user-inputted time to the specified time zone.
func ConvertTime() {
    location := getTimeZone()

    StopDisplayTime() // Stop dynamic display while waiting for user input

    fmt.Println("Enter the time to convert (YYYY-MM-DD HH:mm:ss):")
    var inputTimeStr string
    fmt.Scanln(&inputTimeStr)

    convertedTime, err := time.Parse("2006-01-02 15:04:05", inputTimeStr)
    if err != nil {
        fmt.Println("Invalid time format. Please use the format YYYY-MM-DD HH:mm:ss")
        return
    }

    convertedTime = convertedTime.In(location)
    fmt.Printf("Converted time in %s: %s\n", location, convertedTime.Format("2006-01-02 15:04:05"))
}

func getTimeZone() *time.Location {
    fmt.Println("Select a time zone:")
    zone, err := promptTimeZone()
    if err != nil {
        fmt.Println("Error selecting time zone:", err)
        // Default to UTC if an error occurs
        return time.UTC
    }

    location, err := time.LoadLocation(zone)
    if err != nil {
        fmt.Println("Error loading time zone:", err)
        // Default to UTC if an error occurs
        return time.UTC
    }

    return location
}

func promptTimeZone() (string, error) {
    timeZones := []string{
        "UTC", "America/New_York", "America/Los_Angeles", "Europe/London", "Asia/Tokyo",
        // Add more time zones as needed
    }

    prompt := promptui.Select{
        Label: "Select Time Zone",
        Items: timeZones,
    }

    _, result, err := prompt.Run()
    return result, err
}
