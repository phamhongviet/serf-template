package main

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"os"
	"os/exec"
	//"strconv"
	"log"
	"strings"
	//"text/template"
)

// exit codes
const (
	OK                    = iota
	SYNTAX_ERROR          = iota
	CMD_FAILED            = iota
	TEMPLATE_PARSE_FAILED = iota
)

const (
	DEBUG      = false
	DEBUG_FILE = "/tmp/serf_template.log"
)

func main() {
	if DEBUG {
		log_file, _ := os.Create(DEBUG_FILE)
		defer log_file.Close()
		log.SetOutput(log_file)
		log.Println("Serf Template starting")
	}

	if len(os.Args) != 2 {
		err := errors.New("No config file")
		panic(err)
	}

	// parse directive from config file
	directives, err := ParseDirectives(os.Args[1])
	if err != nil {
		panic(err)
	}

	if DEBUG {
		log.Printf("directives: %s", directives)
	}

	// construct serf command from directive
	cmd_name, cmd_args, err := ConstructSerfCommand(directives)
	if err != nil {
		panic(err)
	}

	if DEBUG {
		log.Printf("CMD: %s %s", cmd_name, cmd_args)
	}

	// exec serf command
	cmd := exec.Command(cmd_name, cmd_args...)
	members_json, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	// parse members
	members, err := ParseMembers(members_json)
	if err != nil {
		panic(err)
	}

	// for each templates:
	// - render template
	// - execute command if any
	for i := 0; i < len(directives.Templates); i++ {
		RenderTemplate(directives.Templates[i].Src, directives.Templates[i].Dest, members)

		if directives.Templates[i].Cmd != "" {
			cmd2_args := strings.Split(directives.Templates[i].Cmd, " ")
			cmd2 := exec.Command(cmd2_args[0], cmd2_args[1:]...)
			err := cmd2.Run()
			if err != nil {
				panic(err)
			}
		}
	}
}
