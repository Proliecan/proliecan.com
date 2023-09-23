package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

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
