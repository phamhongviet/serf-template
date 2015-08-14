package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Template struct {
	Src  string
	Dest string
	Cmd  string
}

type Directive struct {
	Serf        string
	Name        string
	Role        string
	Status      string
	Tags        map[string]string
	Rpc_addr    string        `json:"rpc-addr"`
	Rpc_auth    string        `json:"rpc-auth"`
	Rpc_timeout time.Duration `json:"rpc-timeout"`
	Workers     int           `json:"rpc-workers"`
	Templates   []Template
}

func ParseDirectives(config_file string) (*Directive, error) {
	config_json, err := ioutil.ReadFile(config_file)
	if err != nil {
		return nil, err
	}
	var directive Directive
	err = json.Unmarshal(config_json, &directive)
	if err != nil {
		return nil, err
	}
	// default RPC workers
	if directive.Workers == 0 {
		directive.Workers = 1
	}
	// default RPC address
	if directive.Rpc_addr == "" {
		directive.Rpc_addr = "127.0.0.1:7373"
	}
	// timeout in millisecond. time.Duration use nanosecond by default
	directive.Rpc_timeout = directive.Rpc_timeout * 1000000
	return &directive, nil
}
