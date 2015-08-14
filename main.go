package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

const (
	DEBUG       = false
	DEBUG_FILE  = "/var/log/serf_template.log"
	CONFIG_FILE = "/etc/serf_template/config.json"
)

func main() {
	debug := flag.Bool("debug", DEBUG, "enable debug")
	debugFile := flag.String("log", DEBUG_FILE, "log file for debuging")
	configFile := flag.String("config", CONFIG_FILE, "path to the config file")
	flag.Parse()

	if *debug {
		log_file, _ := os.Create(*debugFile)
		defer log_file.Close()
		log.SetOutput(log_file)
		log.Println("Serf Template starting")
	}

	if *configFile == "" {
		err := errors.New("No config file")
		panic(err)
	}

	// parse directive from config file
	directives, err := ParseDirectives(*configFile)
	if err != nil {
		panic(err)
	}

	if *debug {
		log.Printf("directives: %v", directives)
	}

	agent := NewAgent(&directives)

	agent.run()
}
