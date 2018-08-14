# [Exercism.io - Go](https://exercism.io/tracks/go)

Exercism is an online platform designed to help you improve your coding skills through practice and mentorship.

## List of "Go" Exercises

{{ range . }}
1. {{if .Ready -}} [{{ .Name }}]({{ .Slug }}) {{- else}} {{ .Name }}  {{- end}} {{if .Problems}} {{ range .Problems }}
	- {{if .Ready -}}[{{ .Name }}]({{ .Slug }}){{- else}}{{ .Name }}{{- end}}{{end}}{{- end}}
{{ end }}
