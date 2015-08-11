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
	Tags      map[string]string
	Rpc_addr  string `json:"rpc-addr"`
	Rpc_auth  string `json:"rpc-auth"`
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
	if directive.Rpc_addr == "" {
		directive.Rpc_addr = "127.0.0.1:7373"
	}
	return directive, nil
}
