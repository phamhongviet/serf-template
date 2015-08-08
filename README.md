# serf-template
Template rendering with Serf

## usage

Use as one time runner:

serf-template '/path/to/template-1:/path/to/file-1:command 1' '/path/to/template-2:/path/to/file-2:command 2' ...

Use as Serf handler:

serf agent -event-handler serf-template '/path/to/template-1:/path/to/file-1:command 1'
