package main

import "fmt"

type MyStruct struct {
	string_field string
	int_field int
}

func change_struct(ms *MyStruct) {
	ms.string_field = "bsa"
}

func main() {
	mystruct := &MyStruct{string_field: "abc", int_field: 10}
	change_struct(mystruct) 
	fmt.Println(*mystruct)

}
