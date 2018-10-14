# [LeetCode.com](https://leetcode.com) (Free Exercises only)

The purpose of LeetCode is to provide you hands-on training on real coding interview questions. The Online Judge gives you immediate feedback on the correctness and efficiency of your algorithm which facilitates a great learning experience. Solutions to problem implemented in Go, as I am trying to have more practice in this language.

 * [Progress](#progress)    
 * [Stats](#stats)
 * [Solved Problems](#solved-problemsupcoming-problems)
 * [Upcoming Problems](#upcoming-problems)
 

### Progress  

{{ .Progress }}

### Stats

Dificulty                           | Total | Done
-------------------------------------|-------|------------------ 
☆                                   | {{index .Total "☆"}}   | {{index .Stat "☆"}} 
☆☆                                 | {{index .Total "☆☆"}}  | {{index .Stat "☆☆"}} 
☆☆☆                                | {{index .Total "☆☆☆"}}  | {{index .Stat "☆☆☆"}} 
Total Questions Available            | {{index .Total "All"}}   | {{index .Stat "All"}} 


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