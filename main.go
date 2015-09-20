package main

import (
	"os"
	"flag"
	"github.com/dgarcia202/prunner/core"
)

var concise, export *bool
var source *string

func parseFlags() {
	concise = flag.Bool("concise", false, "Only output result of API calls.")
	export = flag.Bool("export", false, "Save the contents of a url source to a local file named 'source.json'.")
	source = flag.String("source", "", "Required. Specifies the source for the collection information, can be either a postman url or a file path+name.")
	flag.Parse()	
}

func main() {
	
	parseFlags()
	
	runner := core.NewPostmanRunner()
	
	var result bool
	
	if *export {
		result = runner.Export(*source)
	} else {
		result = runner.Run(*source, *concise)
	}
	
	if result {
		os.Exit(0)
	}
	os.Exit(1)	
}