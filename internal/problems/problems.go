package problems

import "sort"

type Problems []Problem

func ConvertMap(problems map[string]Problem) Problems {
	tmp := make(Problems, 0, len(problems))
	for _, problem := range problems {
		tmp = append(tmp, problem)
	}
	sort.Sort(tmp)

	return tmp
}

// Len is part of sort.Interface.
func (p Problems) Len() int {
	return len(p)
}

// Swap is part of sort.Interface.
func (p Problems) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (p Problems) Less(i, j int) bool {
	return p[i].ID > p[j].ID
}
