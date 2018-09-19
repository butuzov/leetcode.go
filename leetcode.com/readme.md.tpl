# [LeetCode.com](https://leetcode.com)
 
 
### Solved Problems

Problem | Dificulty  
--------|----------- 
{{ range . }} 
{{- if .Ready -}} [ {{ .Title }} ]( {{ .Slug }} ) | {{ .LevelStr }}  
{{end}}
{{- end}}

### Upcoming 

Problem @ LeetCode    | Dificulty
----------------------|------------  
{{ range . }} 
{{- if not .Ready -}}[{{ .Title }}](https://leetcode.com/problems/{{- .Slug -}}) | {{ .LevelStr }} 
{{end}}
{{- end}}