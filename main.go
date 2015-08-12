package main

import (
	"errors"
	"log"
	"fmt"
	"os"
)

const (
	DEBUG      = true
	DEBUG_FILE = "/tmp/serf_template.log"
)

func main() {
	fmt.Println("starting")
	// if DEBUG {
	// 	log_file, _ := os.Create(DEBUG_FILE)
	// 	defer log_file.Close()
	// 	log.SetOutput(log_file)
	// 	log.Println("Serf Template starting")
	// }

	if len(os.Args) != 2 {
		err := errors.New("No config file")
		panic(err)
	}

	// parse directive from config file
	directives, err := ParseDirectives(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("parsed directives: %v\n", directives)
	if DEBUG {
		log.Printf("directives: %v", directives)
	}

	agent := NewAgent(&directives)

	agent.run()
}
