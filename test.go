package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(IntArray([]int{10, 20, 30}).ZipAllStringArray(StringArray([]string{"a", "b", "c", "d"}), -100, "NONE").ToString())
}
