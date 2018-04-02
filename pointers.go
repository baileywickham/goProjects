package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*ival = 0
}

func main() {
	i := 1
	fmt.Println("initally", i)
	zeroval(i)
	zeroptr(&i)

}
