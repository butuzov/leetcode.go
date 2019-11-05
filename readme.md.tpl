# [LeetCode.com](https://leetcode.com) (Free Exercises only)

The purpose of LeetCode is to provide you hands-on training on real coding interview questions. The Online Judge gives you immediate feedback on the correctness and efficiency of your algorithm which facilitates a great learning experience. Solutions to problem implemented in Go, as I am trying to have more practice in this language.

 * [Progress](#progress)
 * [Stats](#stats)
 * [Solved Hard Problems](#solved---hard-problems)
 * [Solved Normal Problems](#solved--normal-problems)
 * [Solved Easy Problems](#solved--easy-problems)
 * [Upcoming Problems](#solved-problems)


## Progress

{{ .Progress }}

## Stats

Dificulty                           | Total                     | Done
------------------------------------|---------------------------|------
Easy                                | {{index .Total "Easy"}}   | {{index .Stat "Easy"}}
Normal                              | {{index .Total "Normal"}} | {{index .Stat "Normal"}}
Hard                                | {{index .Total "Hard"}}   | {{index .Stat "Hard"}}
Total Questions Available           | {{index .Total "All"}}    | {{index .Stat "All"}}




## Solved  ☆☆☆ Hard Problems

 ID  | Problem               |  Topics
-----|-----------------------|-------------
{{ range .List }}
{{- if .Hard -}} {{ .ID | printf "%4d" }} | [{{ .Title }}]({{ .Slug }}) | {{ .Topic }}
{{end}}
{{- end}}

### Solved ☆☆ Normal Problems

 ID  | Problem               |  Topics
-----|-----------------------|-------------
{{ range .List }}
{{- if .Normal -}} {{ .ID | printf "%4d" }} | [{{ .Title }}]({{ .Slug }}) | {{ .Topic }}
{{end}}
{{- end}}

## Solved ☆ Easy Problems

 ID  | Problem               | Topics
-----|-----------------------|-------------
{{ range .List }}
{{- if .Easy -}} {{ .ID | printf "%4d" }} | [{{ .Title }}]({{ .Slug }}) | {{ .Topic }}
{{end}}
{{- end}}

## Upsolved Problems

 ID  | Problem @ LeetCode    | Dificulty  | Topics
-----|-----------------------|------------|--------
{{ range .List }}
{{- if not .Ready -}} {{ .ID | printf "%4d" }} | [{{ .Title }}](https://leetcode.com/problems/{{- .Slug -}}) | {{ .LevelStr }}  | {{ .Topic }}
{{end}}
{{- end}}
