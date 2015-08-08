package main

import (
	"errors"
	"strings"
)

type Directive struct {
	template string
	result   string
	command  string
	tags     []string
}

func ParseDirectives(args []string) ([]Directive, error) {
	args_len := len(args)
	directives := make([]Directive, args_len)
	// for each args, parse into directives
	for i := 0; i < args_len; i = i + 1 {
		// split it into parts
		// 1st part: path to template file
		// 2nd part: path to result file
		// 3nd part: command to execute, optional
		// remaining parts: filter tags
		parts := strings.Split(args[i], ":")
		parts_len := len(parts)
		// check number
		if parts_len < 2 {
			return nil, errors.New("Syntax error")
		}
		for i := 0; i < parts_len; i = i + 1 {
			if len(parts[i]) == 0 {
				return nil, errors.New("Syntax error")
			}
		}
		// register directive
		directives[i] = Directive{
			template: parts[0],
			result:   parts[1],
		}
		if parts_len > 2 {
			directives[i].command = parts[2]
		}
		if parts_len > 3 {
			directives[i].tags = make([]string, parts_len-3)
			for j := 0; j < parts_len-3; j = j + 1 {
				directives[i].tags[j] = parts[3+j]
			}
		}
	}

	return directives, nil
}
