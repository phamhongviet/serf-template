package main

import (
	"encoding/json"
	"io/ioutil"
)

type Template struct {
	Src  string
	Dest string
	Cmd  string
}

type Directive struct {
	Serf      string
	Name      string
	Role      string
	Status    string
	Tags      []string
	Rpc_addr  string
	Rpc_auth  string
	Templates []Template
}

func ParseDirectives(config_file string) (Directive, error) {
	config_json, err := ioutil.ReadFile(config_file)
	if err != nil {
		panic(err)
	}
	var directive Directive
	err = json.Unmarshal(config_json, &directive)
	if err != nil {
		panic(err)
	}
	return directive, nil
}
