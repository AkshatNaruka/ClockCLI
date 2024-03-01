package main

import (
	"bufio"
	"clock/pkg"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"
	"github.com/manifoldco/promptui"
	
)
var server *http.Server

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func startClockServer() {
	if server != nil {
		// Server is already running
		return 
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, _ := ioutil.ReadFile("../pkg/clock.html")
		_, _ = w.Write(content)
	})


	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
        stopClockServer()
        w.WriteHeader(http.StatusOK)
        // You can optionally send a response message here if needed
    })

	server = &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s\n", err)
		}
	}()
}

func stopClockServer() {
	if server != nil {
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Fatalf("Server shutdown error: %s\n", err)
		}
	}
	server = nil
}

func main() {
	prompt := promptui.Select{
		Label: "Select Utility",
		Items: []string{"TerminalClock","Clock","Stopwatch", "Alarm","Calendar","Exit"},
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	for {
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "TerminalClock":
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
		case "Calendar":
			pkg.DisplayCalendar()
		case "Exit":
			fmt.Println("Exiting the CLI...")
			os.Exit(0)
		case "Clock":
			startClockServer()
			openBrowser("http://localhost:8080")
			fmt.Println("Press 'q' to quit")
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _, err := reader.ReadRune()
				if err != nil {
					fmt.Println(err)
					break
				}

				if char == 'q' || char == 'Q'{
					stopClockServer()
					break
				}
			}
		default:
			fmt.Println("Please choose a utility to run.")
		}
	}
}
			
		
