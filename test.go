package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(IntArray([]int{10, 20, 30}).ZipAllStringArray(StringArray([]string{"a", "b", "c", "d"}), -100, "NONE").ToString())

	t1 := Tuple2{10, "hello"}
	t2 := Tuple2{t1, MakeIntList(1, 2, 3)}
	fmt.Println(t2.ToString())
}
