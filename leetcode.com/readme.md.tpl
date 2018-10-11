# [LeetCode.com](https://leetcode.com)

 * [Progress](#progress)
 * [Stats](#stats)
 * [Solved Problems](#solved-problems)
 * [Upcoming Problems](#upcoming-problems)

### Progress

{{ .Progress }}

### Stats 

Dificulty                            |Done
-------------------------------------|------------------------- 
☆                                     | {{index .Stat "☆"}} 
☆☆                                   | {{index .Stat "☆☆"}} 
☆☆☆                                 | {{index .Stat "☆☆☆"}} 
Total Questions Available             | {{index .Stat "All"}} 


### Solved Problems

ID | Problem               | Dificulty  | Topics
---|-----------------------|------------|------------ 
{{ range .List }} 
{{- if .Ready -}} {{ .ID }} | [ {{ .Title }} ]( {{ .Slug }} ) | {{ .LevelStr }} | {{ .Topic }} 
{{end}}
{{- end}}

### Upcoming Problems

ID | Problem @ LeetCode    | Dificulty  | Topics
---|-----------------------|------------|--------
{{ range .List }} 
{{- if not .Ready -}}{{ .ID }} | [{{ .Title }}](https://leetcode.com/problems/{{- .Slug -}}) | {{ .LevelStr }}  | {{ .Topic }} 
{{end}}
{{- end}}