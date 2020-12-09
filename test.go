package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l1 := MakeIntList(1, 2, 3)
	l2 := MakeStringList("a", "b", "c")
	l3 := Float32Array([]float32{0.5, 0.7, 0.9})

	fmt.Println(l1.ZipStringList(l2).ZipFloat32Array(l3).ToString())

	fmt.Println(
		MakeIntList(1, 2, 3).
			ZipAllStringList(
				MakeStringList("a", "b", "c", "d"), -100, "NONE").
			ToString())

	fmt.Println(StringArray([]string{"a", "b", "c"}).ZipWithIndex().ToString())
}
