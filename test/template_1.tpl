BEGIN
{{ range . }}{{ if eq .Status "alive" }}server {{.Name}} at {{.Addr}} with serf at {{.Port}} and {{.Tags.app}} at {{.Tags.webport}}
{{ end }}{{ end }}
END
