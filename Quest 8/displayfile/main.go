package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	body, err := ioutil.ReadFile(args[1])
	if err == nil {
		fmt.Print(string(body))
	}
}
