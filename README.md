# Serf Template
Template rendering with Serf.

## Installation

	$ git clone https://github.com/phamhongviet/serf-template.git
	$ cd serf-template
	$ go build

## Usage

	$ serf-template config-file

Instead of reading anything from stdin or environment variables, Serf Template get members' information from the RPC interface. Serf Temlate subscribe to member related events from RPC stream (member-join, member-failed, member-update, member-leave and member-reap).

## Configuration

Example:

	{
	  "serf": "path/to/serf"
	  "name": "regexp",
	  "role": "regexp",
	  "status": "regexp",
	  "tags": {
	    "key1": "value1",
	    "key2": "value2",
	    "key3": "value3"
	  },
	  "rpc-addr": "127.0.0.1:7373",
	  "rpc-auth": "rpcauthtoken",
	  "rpc-timeout": 1000,
	  "templates": [
	    {
	    "src": "/path/to/template.tpl",
	    "dest": "/path/to/result.file",
	    "cmd": "service dummy restart"
	    },
	    {
	    "src": "/path/to/src.file",
	    "dest": "/path/to/dest.file",
	    "cmd": "service dummier restart",
	    }
	  ]
	}

## Template File

Serf Template consumes template files in [Go Template][] format. Template files are rendered with a list of members from executing `serf members -format json` command. A member has `Name`, `Addr`, `Port`, `Tags` and `Status`.

Member structure example:

	{
	  "name": "web-1",
	  "addr": "172.16.0.21",
	  "port": 7946,
	  "tags": {
	    "webport": "8080",
	    "app": "nginx",
	    "role": "web"
	  },
	  "status": "failed",
	}

Template file example:

	{{ range . }}{{ if eq .Status "alive" }}
	server {{.Name}} at {{.Addr}} with serf at {{.Port}} and {{.Tags.app}} at {{.Tags.webport}}
	{{ end }}{{ end }}

The above template file would produce a file like this:

	server web-1 at 172.16.0.21 with serf at 7946 and nginx at 8080
	server web-2 at 172.16.0.22 with serf at 7946 and nginx at 8080
	server web-3 at 172.16.0.23 with serf at 7946 and httpd at 80

[Go Template]: http://golang.org/pkg/text/template/ "Go Template"
