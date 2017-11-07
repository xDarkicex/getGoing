package main

import (
	"flag"
	"log"
	"os"

	logic "github.com/xDarkicex/Logic"
	"github.com/xDarkicex/getGoing/generator"
	"github.com/xDarkicex/getGoing/helpers"
)

// Project version number
const version = "1.0.0"

/*
Default values for generation of new project
defaultApp is the default name for app if no new flag is used
defaultValue is the value of the manual function
*/
const (
	defaultApp      = ""
	defaultValue    = false
	defaultTemplate = "basic"
)

// storing var for export
var (
	template string
	dirPath  string
	appName  string
	man      bool
	m        bool
	Generate = generator.Generate
	Manual   = helpers.Manual
	Xor      = logic.Xor
)

/* func init()
Using the init func to load in all flags before stepping into main func
*/

func init() {
	flag.StringVar(&appName, "new", defaultApp, "generate default application")
	flag.StringVar(&template, "template", defaultTemplate, "template choice for generation")
	flag.BoolVar(&man, "man", defaultValue, "Usage manual")
	flag.BoolVar(&man, "m", defaultValue, "Usage manual")
}

func main() {
	flag.Parse()

	if Xor(man, m) {
		if err := Manual(man, m); err != nil {
			log.Fatal(err)
		}
	}

	if err := Generate(appName, template); err != nil {
		log.Fatal(err)
	}

	// return 0;
	os.Exit(0)
}
