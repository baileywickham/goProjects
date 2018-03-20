package main

import "fmt"

func main() {
	i := 1
	for i < 4 {
		fmt.Println(i)
	i = i + 1
	}
	s := make([]string, 3)
	//creates a empty slice of len 3
	s[0] = "a"
	//Can be treated as a combonation of list and array.
	//Lower level than python, carries verbosity. 
	t := make([]byte, 3, 3)

}

