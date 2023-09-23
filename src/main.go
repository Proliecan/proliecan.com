package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var verbose bool = false

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// parse args
	config := parseArgs()
	if verbose {
		log.Println("Config:")
		log.Println("\tDomain:", config.Domain)
		log.Println("\tPort:", config.Port)
	}

	log.Println("Starting server...")
	initServer()

	// start server
	go startServer(config.Domain, config.Port)

	// print server info
	log.Println("Server started at", colorize(config.Domain, Yellow)+":"+colorize(fmt.Sprintf("%d", config.Port), Yellow))

	<-done
	log.Println("\nServer Stopped gracefully")
}

type Config struct {
	Domain string
	Port   int
}

func parseArgs() Config {
	bin := os.Args[0]

	var config Config

	// for all args
	for i, arg := range os.Args {
		if arg == "-d" {
			if i+1 < len(os.Args) {
				config.Domain = os.Args[i+1]
				continue
			}
		}
		if arg == "-p" {
			if i+1 < len(os.Args) {
				port, err := strconv.Atoi(os.Args[i+1])
				if err != nil {
					log.Println(colorize("Error:", Red), "Invalid port")
					printUsage(bin)
					os.Exit(1)
				}
				config.Port = port
				continue
			}
		}
		if arg == "-v" {
			verbose = true
			continue
		}
		if arg == "-h" {
			printUsage(bin)
			os.Exit(0)
		}
	}

	// check if domain is set
	if config.Domain == "" {
		log.Println(colorize("Error:", Red), "Domain not set")
		printUsage(bin)
		os.Exit(1)
	}

	// check if port is set
	if config.Port == 0 {
		log.Println(colorize("Error:", Red), "Port not set")
		printUsage(bin)
		os.Exit(1)
	}

	return config
}

func printUsage(bin string) {
	log.Println(colorize("Usage:", Red),
		colorize(bin, Yellow), colorize("-d", Cyan), colorize(colorize("<domain>", CyanBG), Red),
		colorize("-p", Green), colorize(colorize("<port>", GreenBG), Red),
		colorize("[-v]", Blue),
		colorize("[-h]", Magenta))

	log.Println(colorize("Example:", Red),
		colorize(bin, Yellow), colorize("-d", Cyan), colorize(colorize("localhost", CyanBG), Red),
		colorize("-p", Green), colorize(colorize("8080", GreenBG), Red),
		colorize("-v", Blue))
}
