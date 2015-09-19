package main

import (
	"fmt"
	"os"
	"github.com/dgarcia202/prunner/core"
)

func main() {
	fmt.Println("Welcome to postman runner!")
	runner := core.NewPostmanRunner()
	runner.Run(os.Args[1])
}