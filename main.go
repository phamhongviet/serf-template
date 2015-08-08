package main

import (
	//"encoding/json"
	//"errors"
	"fmt"
	"os"
	"os/exec"
	//"strconv"
	"strings"
	"text/template"
)

// exit codes
const (
	OK                    = iota
	SYNTAX_ERROR          = iota
	CMD_FAILED            = iota
	TEMPLATE_PARSE_FAILED = iota
)

type Member struct {
	Name string
	Addr string
	Role string
	Tags map[string]string
}

func main() {
	directives, err := ParseDirectives(os.Args[1:])
	if err != nil {
		panic(err)
	}
	directives_len := len(directives)

	var members []Member
	// render template for each directives
	for i := 0; i < directives_len; i = i + 1 {
		// parse template
		tpl, err := template.ParseFiles(directives[i].template)
		if err != nil {
			fmt.Println(err)
			os.Exit(TEMPLATE_PARSE_FAILED)
		}

		// render template
		result_file, err := os.Create(directives[i].result)
		defer result_file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(TEMPLATE_PARSE_FAILED)
		}
		err = tpl.Execute(result_file, members)
		if err != nil {
			fmt.Println(err)
			os.Exit(TEMPLATE_PARSE_FAILED)
		}

		// execute command
		if directives[i].command != "" {
			cmd_args := strings.Split(directives[i].command, " ")
			cmd := exec.Command(cmd_args[0], cmd_args[1:]...)
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(CMD_FAILED)
			}
		}
	}

	os.Exit(OK)
}
