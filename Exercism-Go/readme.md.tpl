# [Exercism.io - Go](https://exercism.io/tracks/go)

Exercism is an online platform designed to help you improve your coding skills through practice and mentorship.

## List of "Go" Exercises 
{{define "ROW"}} {{if .Ready -}}[{{if .Problems}}**{{.Name}}**{{else}}{{ .Name }}{{end}}]({{ .Slug }}){{- else}}{{ .Name }}{{- end}} |  [exercism.io / {{ .Slug }} ](https://exercism.io/tracks/go/exercises/{{ .Slug }})  
{{end}}
 

 Problem @ GitHub Solution     | Problem Page @ Exercism  
-------------------------------|------------------------- 
{{ range . }}{{template "ROW" .}}{{if .Problems}}{{ range .Problems }}{{template "ROW" .}}{{ end }}{{ end }}{{end}}
