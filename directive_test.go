package main

import (
	"os"
	"fmt"
)

func ExampleParseDirectives() {
	var d *Directive
	var e error
	// example 1: bad path
	d, e = ParseDirectives("test/unknown.json")
	if e == nil || d != nil {
		panic(e)
	}

	// example 2: bad json
	if _, e = os.Create("test/empty"); e != nil {
		panic(e)
	}
	d, e = ParseDirectives("test/empty")
	if e == nil || d != nil {
		panic(e)
	}
	if e = os.Remove("test/empty"); e != nil {
		panic(e)
	}

	// example 3: no config
	d, e = ParseDirectives("test/config_1.json")
	if e != nil {
		panic(e)
	}
	fmt.Println(*d)

	// example 4: full config
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
	fmt.Println(d.Workers)
	fmt.Println(d.Templates[0].Src)
	fmt.Println(d.Templates[0].Dest)
	fmt.Println(d.Templates[0].Cmd)
	// Output:
	// {    map[] 127.0.0.1:7373  0 1 []}
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
	// 5
	// /path/to/template.tpl
	// /path/to/result.file
	// ls
}
