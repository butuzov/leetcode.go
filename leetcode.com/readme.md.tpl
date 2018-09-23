# [LeetCode.com](https://leetcode.com)
 
 
### Solved Problems

ID | Problem               | Dificulty  | Topics
---|-----------------------|------------|------------ 
{{ range . }} 
{{- if .Ready -}} {{ .ID }} | [ {{ .Title }} ]( {{ .Slug }} ) | {{ .LevelStr }} | {{ .Topic }} 
{{end}}
{{- end}}

### Upcoming 

ID | Problem @ LeetCode    | Dificulty  | Topics
---|-----------------------|------------|--------
{{ range . }} 
{{- if not .Ready -}}{{ .ID }} | [{{ .Title }}](https://leetcode.com/problems/{{- .Slug -}}) | {{ .LevelStr }}  | {{ .Topic }} 
{{end}}
{{- end}}