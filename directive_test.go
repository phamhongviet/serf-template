package main

import (
	//"testing"
	"fmt"
)

func ExampleParseDirectives() {
	var d Directive
	var e error
	// example 1: no config
	d, e = ParseDirectives("test/config_1.json")
	if e != nil {
		panic(e)
	}
	fmt.Println(d)

	// example 2: full config
	d, e = ParseDirectives("test/config_2.json")
	if e != nil {
		panic(e)
	}
	fmt.Println(d.Serf)
	fmt.Println(d.Name)
	fmt.Println(d.Role)
	fmt.Println(d.Status)
	fmt.Println(len(d.Tags))
	fmt.Println(d.Tags[0])
	fmt.Println(d.Tags[1])
	fmt.Println(d.Rpc_addr)
	fmt.Println(d.Rpc_auth)
	fmt.Println(d.Templates[0].Src)
	fmt.Println(d.Templates[0].Dest)
	fmt.Println(d.Templates[0].Cmd)
	// Output:
	// {    []   []}
	// serf
	// svr
	// web
	// alive
	// 2
	// app=a1
	// port=80
	// 127.0.0.1:7373
	// rpcauthtoken
	// /path/to/template.tpl
	// /path/to/result.file
	// service dummy restart
}
