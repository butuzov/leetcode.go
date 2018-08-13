package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldl- list itteration.
func (l IntList) Foldl(f binFunc, initial int) int {

	for _, i := range l {
		initial = f(initial, i)
	}

	return initial
}

// Foldr - reverced list itteration.
func (l IntList) Foldr(f binFunc, initial int) int {

	for i := len(l) - 1; i >= 0; i-- {
		initial = f(l[i], initial)
	}

	return initial
}

// Filter based on predicate.
func (l IntList) Filter(f predFunc) IntList {
	var nl IntList = []int{}
	for _, i := range l {
		if f(i) {
			nl = append(nl, i)
		}
	}
	return nl
}

// Length return length of list.
func (l IntList) Length() int {
	var n int
	for _ = range l {
		n += 1
	}
	return n
}

// Map - applying unary function to all elements of list.
func (l IntList) Map(f unaryFunc) IntList {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}

// Reverse list.
func (l IntList) Reverse() IntList {

	for i, n := 0, l.Length()-1; i < n; i, n = i+1, n-1 {
		l[i], l[n] = l[n], l[i]
	}

	return l
}

// Append element oflist a to list l.
func (l IntList) Append(a IntList) IntList {

	for _, i := range a {
		l = append(l, i)
	}

	return l
}

// Concat elements of element of lists to initial list l.
func (l IntList) Concat(lists []IntList) IntList {

	for _, list := range lists {
		for _, i := range list {
			l = append(l, i)
		}
	}

	return l
}
