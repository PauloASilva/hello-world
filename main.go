package main

import (
	"flag"
	"github.com/PauloASilva/hello-world/config"
	"github.com/PauloASilva/hello-world/web"
)

// Application command line flags
type flags struct {
	ConfigFile string
}

func main() {
	var f flags
	var c config.Config

	// handle command line flags
	flag.StringVar(&f.ConfigFile, "c", "", "Configuration file")
	flag.Parse()

	// load configuration file
	c.Load(f.ConfigFile)

	// create the server
	web.NewServer(&c)
}
