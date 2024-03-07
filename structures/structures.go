package main

import "fmt"

func main() {
	b := [2]int{1, 2}
	change_array(&b)
	fmt.Println(b)
}
