package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(IntArray([]int{10, 20, 30}).ZipStringArray(StringArray([]string{"a"})).ToString())
}
