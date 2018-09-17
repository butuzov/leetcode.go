# [Exercism.io - Go](https://exercism.io/tracks/go)

Exercism is an online platform designed to help you improve your coding skills through practice and mentorship.

### Solved Problems

Problem | Dificulty | Tags
--------|-----------|-------------------------------------------------
{{ range . }} 
{{- if .Ready -}} [{{ .Title }} ]( {{- .Alias -}} ) | {{ .Level }} | {{ .Topic }} 
{{end}}
{{- end}}

### Upcoming 

Problem @ Exercsim.io | Dificulty | Tags
----------------------|-----------|-------------------------------------------------
{{ range . }} 
{{- if .Ready -}}{{else}} [{{ .Title }} @ exercism.io ](https://exercism.io/tracks/go/exercises/{{- .Alias -}}) | {{ .Level }} | {{ .Topic }} 
{{end}}
{{- end}}