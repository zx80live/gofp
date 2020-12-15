package main

import . "github.com/zx80live/gofp/fp"
import "fmt"

func main() {
	it := IntOpt(100).Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}

}
