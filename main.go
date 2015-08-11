package main

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"os"
	"os/exec"
	//"strconv"
	rpc "github.com/hashicorp/serf/client"
	"log"
	"strings"
	//"text/template"
)

// exit codes
/*
const (
	OK                    = iota
	SYNTAX_ERROR          = iota
	CMD_FAILED            = iota
	TEMPLATE_PARSE_FAILED = iota
)
*/

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

	// create RPC config
	rpc_config := rpc.Config{
		Addr:    directives.Rpc_addr,
		AuthKey: directives.Rpc_auth,
		Timeout: directives.Rpc_timeout,
	}
	// create connection to the RPC interface
	rpc_client, err := rpc.ClientFromConfig(&rpc_config)
	if err != nil {
		panic(err)
	}

	// get members' information
	members, err := rpc_client.MembersFiltered(directives.Tags, directives.Status, directives.Name)
	if err != nil {
		panic(err)
	}

	// for each templates:
	// - render template
	// - execute command if any
	for i := 0; i < len(directives.Templates); i++ {
		RenderTemplate(directives.Templates[i].Src, directives.Templates[i].Dest, members)

		if directives.Templates[i].Cmd != "" {
			cmd_args := strings.Split(directives.Templates[i].Cmd, " ")
			cmd := exec.Command(cmd_args[0], cmd_args[1:]...)
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}
	}
}
