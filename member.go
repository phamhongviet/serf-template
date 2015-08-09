package main

import "strings"
import "errors"

type Member struct {
	Name     string
	Addr     string
	Port     int
	Tags     map[string]string
	Status   string
	Protocol map[string]int
}

type SerfOutput struct {
	Members []Member
}
