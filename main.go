package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // new pointer to variable a

	// obtaining a pointer to variable of type int
	// init
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c and a unchanged

	c = d   // c and d point to the same location
	*c = 14 // с = 14 -> d = 14, a = 12
}
