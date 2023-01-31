package main

import "fmt"

func main() {
	a := make([]byte, 26)
	a[0] = 1
	a[0]--
	fmt.Println(len(a))
}
