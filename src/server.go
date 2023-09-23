package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	yaml "gopkg.in/yaml.v3"
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

type Config struct {
	Domain string
	Port   int
}

func (c *Config) readFromFile(path string) {
	if verbose {
		log.Println("Reading config file from", path)
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if verbose {
		fmt.Printf("Domain: %s\n", c.Domain)
		fmt.Printf("Port: %d\n", c.Port)
	}
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
