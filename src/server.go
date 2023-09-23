package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting server...")
	initServer()

	// start server
	go func() {
		log.Println("Server Started")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	<-done
	log.Println("\nServer Stopped gracefully")
}

func initServer() {
	// handle /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(colorize(r.RemoteAddr, Yellow),
			colorize(">", Red), r.Method, r.URL.Path)
		log.Println(colorize(r.RemoteAddr, Yellow),
			colorize("<", Green), "Status:", http.StatusOK)
		w.WriteHeader(http.StatusOK)
	})
}

// colorize string
type ColorCode string

const (
	Reset     ColorCode = "\033[0m"
	Red       ColorCode = "\033[31m"
	Green     ColorCode = "\033[32m"
	Yellow    ColorCode = "\033[33m"
	Blue      ColorCode = "\033[34m"
	Magenta   ColorCode = "\033[35m"
	Cyan      ColorCode = "\033[36m"
	White     ColorCode = "\033[37m"
	RedBG     ColorCode = "\033[41m"
	GreenBG   ColorCode = "\033[42m"
	YellowBG  ColorCode = "\033[43m"
	BlueBG    ColorCode = "\033[44m"
	MagentaBG ColorCode = "\033[45m"
	CyanBG    ColorCode = "\033[46m"
	WhiteBG   ColorCode = "\033[47m"
	Bold      ColorCode = "\033[1m"
	Dim       ColorCode = "\033[2m"
	Italic    ColorCode = "\033[3m"
	Underline ColorCode = "\033[4m"
)

func colorize(s string, color ColorCode) string {
	return string(color) + s + string(Reset)
}
