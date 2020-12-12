package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileLinesList(path string) {
	var file, err = os.Open("/path/to/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
}

func WithFile(path string) func(func(*os.File)) {
	return func(f func(*os.File)) {
		var file, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		f(file)
	}
}

func WithScan(path string) func(func(*bufio.Scanner)) {
	return func(f func(*bufio.Scanner)) {
		var file, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		s := bufio.NewScanner(file)
		f(s)
	}
}

func main() {
	fmt.Println("fp.io")
	WithFile("/home/temp/tmp.txt")(func(file *os.File) {
		s := bufio.NewScanner(file)
		for s.Scan() {
			fmt.Println(s.Text())
		}
	})

	//var file, err = os.Open("/path/to/file.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
}
