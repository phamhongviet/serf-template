package main

import (
	//"fmt"
	"os"
	//"os/exec"
	"strings"
)

// exit codes
const (
	OK           = iota
	SYNTAX_ERROR = iota
)

type Directive struct {
	template string
	result   string
	command  string
	tags     []string
}

func main() {
	args_len := len(os.Args)
	directives := make([]Directive, args_len-1)
	// for each args
	for i := 1; i < args_len; i = i + 1 {
		// split it into parts
		// 1st part: path to template file
		// 2nd part: path to result file
		// 3nd part: command to execute, optional
		// remaining parts: filter tags
		parts := strings.Split(os.Args[i], ":")
		parts_len := len(parts)
		// check number
		if parts_len < 2 {
			os.Exit(SYNTAX_ERROR)
		}
		for i := 0; i < parts_len; i = i + 1 {
			if len(parts[i]) == 0 {
				os.Exit(SYNTAX_ERROR)
			}
		}
		// register directive
		directives[i-1] = Directive{
			template: parts[0],
			result:   parts[1],
		}
		if parts_len > 2 {
			directives[i-1].command = parts[2]
		}
		if parts_len > 3 {
			directives[i-1].tags = make([]string, parts_len-3)
			for j := 0; j < parts_len-3; j = j + 1 {
				directives[i-1].tags[j] = parts[2+j]
			}
		}
	}

	os.Exit(OK)
}
