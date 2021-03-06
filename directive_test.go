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
	fmt.Println(d.Tags["app"])
	fmt.Println(d.Tags["port"])
	fmt.Println(d.Rpc_addr)
	fmt.Println(d.Rpc_auth)
	fmt.Println(d.Rpc_timeout)
	fmt.Println(d.Templates[0].Src)
	fmt.Println(d.Templates[0].Dest)
	fmt.Println(d.Templates[0].Cmd)
	// Output:
	// {    map[] 127.0.0.1:7373  0 []}
	// serf
	// svr
	// web
	// alive
	// 2
	// a1
	// 80
	// 127.0.0.1:7373
	// rpcauthtoken
	// 500ms
	// /path/to/template.tpl
	// /path/to/result.file
	// service dummy restart
}
