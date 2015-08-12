package main

import (
	rpc "github.com/hashicorp/serf/client"
)

type Agent struct {
	Directives *Directive
	RPCConfig  *rpc.Config
	DeadNode   chan bool
}

func NewAgent(directives *Directive) Agent {
	// create RPC config
	rpc_config := rpc.Config{
		Addr:    directives.Rpc_addr,
		AuthKey: directives.Rpc_auth,
		Timeout: directives.Rpc_timeout,
	}

	deadNode := make(chan bool)
	for i := 0; i < directives.Workers; i++ {
		w := NewWorker(directives, &rpc_config, deadNode)
		w.run()
	}
	return Agent{
		Directives: directives,
		RPCConfig:  &rpc_config,
		DeadNode:   deadNode,
	}
}

func (a *Agent) run() {
	for {
		_ = <-a.DeadNode
		w := NewWorker(a.Directives, a.RPCConfig, a.DeadNode)
		w.run()
	}
}
