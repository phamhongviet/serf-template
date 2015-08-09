package main

import (
	"io/ioutil"
	"encoding/json"
)

type Template struct {
	Src string
	Dest string
	Cmd string
}

type Directive struct {
	Serf string
	Name string
	Role string
	Status string
	Tags []string
	Rpc-addr string
	Rpc-auth string
	Templates []Template
}

func ParseDirectives(config-file string) (Directive, error) {
	config-json, err := ioutil.ReadFile(config-file)
	if err != nil {
		panic(err)
	}
	var directive Directive
	err = json.Unmarshal(config-json, &directive)
	if err != nil {
		panic(err)
	}
	return directive, nil
}
