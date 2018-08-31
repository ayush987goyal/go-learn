package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
