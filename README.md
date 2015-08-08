# Serf Template
Template rendering with Serf

## Installation

	$ git clone https://github.com/phamhongviet/serf-template.git
	$ cd serf-template
	$ go build

## Usage

Use as one time runner:

	serf-template directive-1 directive-2 directive-3 ...

Use as Serf handler:

	serf agent -event-handler serf-template serf-template directive-1 directive-2 directive-3 ...

## Directive format

	template:result:command:tag-1:tag-2:...:tag-n

Example:

	/etc/haproxy/haproxy.tpl:/etc/haproxy/haproxy.cfg:service haproxy restart:role=web
	/etc/varnish/varnish.tpl:/etc/varnish/default.vcl:service varnish reload:role=web:app=image-processor:dc=us-west

## Template files

Serf Template consumes template files in [Go Template][] format. Template files are rendered with a list of members from executing `serf members -format json` command. A member has `Name`, `Addr`, `Port`, `Tags`, `Status` and `Protocol`.

Example:

	{{ range . }}{{ if eq .Status "alive" }}
	server {{.Name}} at {{.Addr}} with {{.Port}}
	{{ end }}{{ end }}

The above template file would produce a file like this:

	server web-1 at 172.17.0.21 with 7946
	server web-2 at 172.17.0.22 with 7946
	server web-3 at 172.17.0.23 with 7946

[Go Template]: http://golang.org/pkg/text/template/ "Go Template"
