// prunner.go

/*
* Usage: prunner file|url <flags>
*
* 	-debug		log file with debug info is generated
*/

package main

import (
	"flag"
)

type Config struct {
	source string
	debug bool
}

var config Config

func readConfig() {
	
	flag.Parse()
}

func main() {
	readConfig()
}