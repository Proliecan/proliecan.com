package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var verbose bool = false

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// read config file path from command line
	var configPath string = parseArgs()
	// read config file
	config := Config{}
	config.readFromFile(configPath)

	log.Println("Starting server...")
	initServer()

	// start server
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Domain, config.Port), nil))
	}()
	// print server info
	log.Println("Server started at", colorize(config.Domain, Yellow)+":"+colorize(fmt.Sprintf("%d", config.Port), Yellow))

	<-done
	log.Println("\nServer Stopped gracefully")
}

func parseArgs() string {
	bin := os.Args[0]
	// find pos of -c
	var configPath string
	for i, arg := range os.Args {
		if arg == "-c" {
			if i+1 < len(os.Args) {
				configPath = os.Args[i+1]
				continue
			}
		}
		if arg == "-v" {
			verbose = true
			continue
		}
	}

	if configPath == "" {
		printUsage(bin)
		os.Exit(1)
	}

	return configPath
}

func printUsage(bin string) {
	log.Println(colorize("Usage:", Red),
		colorize(bin, Yellow), colorize("-c", Cyan), colorize("<path to config file>", Cyan),
		colorize("[-v]", Blue))
}
