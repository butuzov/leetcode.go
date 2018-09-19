// example of the global (leacked variable and one that dies at
// end of the scope)
package main

var global *int

func main() {
	f()
	g()
}

func f() {

	var x int
	x = 1
	global = &x

}
func g() {
	y := new(int)
	*y = 1
}
