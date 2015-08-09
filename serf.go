package main

func ConstructSerfCommand(drt Directive) (name string, args []string, err error) {
	err = nil
	if drt.Serf != "" {
		name = drt.Serf
	} else {
		// assume serf in PATH
		name = "serf"
	}
	// TODO: check if `name` is executable
	args = []string{"members", "-format", "json"}
	if drt.Name != "" {
		args = append(args, "-name", drt.Name)
	}
	if drt.Role != "" {
		args = append(args, "-role", drt.Role)
	}
	if drt.Status != "" {
		args = append(args, "-status", drt.Status)
	}
	for i := 0; i < len(drt.Tags); i++ {
		args = append(args, "-tag", drt.Tags[i])
	}
	if drt.Rpc_addr != "" {
		args = append(args, "-rpc-addr", drt.Rpc_addr)
	}
	if drt.Rpc_auth != "" {
		args = append(args, "-rpc-auth", drt.Rpc_auth)
	}
	return
}
