package main

import "fmt"

func ExampleConstructSerfCommand() {
	d, e := ParseDirectives("test/config_2.json")
	if e != nil {
		panic(e)
	}
	name, args, e := ConstructSerfCommand(d)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s %s", name, args)
	// Output:
	// serf [members -format json -name svr -status alive -tag role=web -tag app=a1 -tag port=80 -rpc-addr 127.0.0.1:7373 -rpc-auth rpcauthtoken]
}
