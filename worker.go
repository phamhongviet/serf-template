package main

import (
	rpc "github.com/hashicorp/serf/client"
	"os/exec"
	"strings"
)

type Worker struct {
	Directives *Directive
	Client     *rpc.RPCClient
	DeadNode   chan bool
}

func NewWorker(config *Directive, rpc_config *rpc.Config, deadNode chan bool) *Worker {
	// create connection to the RPC interface
	rpc_client, err := rpc.ClientFromConfig(rpc_config)
	if err != nil {
		panic(err)
	}

	w := Worker{
		Directives: config,
		Client:     rpc_client,
		DeadNode:   deadNode,
	}
	go w.run()

	return &w
}

func (w *Worker) run() {
	defer func() {
		w.Client.Close()
		w.DeadNode <- true
	}()

	ch := make(chan map[string]interface{})
	suscription, err := w.Client.Stream("member-join,member-failed,member-update,member-leave,member-reap", ch)
	if err != nil {
		panic(err)
	}
	defer w.Client.Stop(suscription)

	for {
		// wait for signal from serf
		<-ch

		if err = w.processTemplates(); err != nil {
			panic(err)
		}
	}
}

func (w *Worker) processTemplates() error {
	members, err := w.Client.MembersFiltered(w.Directives.Tags, w.Directives.Status, w.Directives.Name)
	if err != nil {
		return err
	}

	// for each templates:
	// - render template
	// - execute command if any
	for i := 0; i < len(w.Directives.Templates); i++ {
		RenderTemplate(w.Directives.Templates[i].Src, w.Directives.Templates[i].Dest, members)

		if w.Directives.Templates[i].Cmd != "" {
			err = w.runCmd(w.Directives.Templates[i].Cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Worker) runCmd(cmdString string) error {
	cmd_args := strings.Split(cmdString, " ")
	cmd := exec.Command(cmd_args[0], cmd_args[1:]...)
	return cmd.Run()
}
