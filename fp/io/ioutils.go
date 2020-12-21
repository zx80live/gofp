package io

import (
	"bufio"
	"github.com/zx80live/gofp/fp"
	"log"
	"os"
)

func FileStringLines(path string) fp.StringQueue {
	var file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	acc := fp.MkStringQueue()
	for scanner.Scan() {
		acc = acc.Enqueue(scanner.Text())
	}
	return acc
}
