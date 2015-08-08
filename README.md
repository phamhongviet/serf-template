# serf-template
Template rendering with Serf

## usage

Use as one time runner:

serf-template directive-1 directive-2 directive-3 ...

Use as Serf handler:

serf agent -event-handler serf-template serf-template directive-1 directive-2 directive-3 ...

## directive format

template:result:command:tag-1:tag-2:...:tag-n

Example:

	/etc/haproxy/haproxy.tpl:/etc/haproxy/haproxy.cfg:service haproxy restart:role=web
	/etc/varnish/varnish.tpl:/etc/varnish/default.vcl:service varnish reload:role=web:app=image-processor:dc=us-west
