package main

import (
	"fmt"
	"os"
	"github.com/dgarcia202/prunner/core"
)

func main() {
	fmt.Println("Welcome to postman runner!")
	runner := core.NewPostmanRunner()
	result := runner.Run(os.Args[1])
	
	if result {
		os.Exit(0)
	}
	os.Exit(1)	
}