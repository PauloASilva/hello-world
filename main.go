package main

import (
	"flag"
	"github.com/PauloASilva/hello-world/config"
	"github.com/PauloASilva/hello-world/web"
	"github.com/PauloASilva/m1"
	"github.com/PauloASilva/m2"
	"github.com/PauloASilva/m3"
	"github.com/PauloASilva/m4"
)

// Application command line flags
type flags struct {
	ConfigFile string
}

func main() {
	var f flags
	var c config.Config
	var s web.Server

	// handle command line flags
	flag.StringVar(&f.ConfigFile, "c", "", "Configuration file")
	flag.Parse()

	// load configuration file
	c.Load(f.ConfigFile)

	// create the server
	s = web.CreateServer(&c)

	// register modules
	s.RegisterModule("/m1", m1.Routes)
	s.RegisterModule("/m2", m2.Routes)
	s.RegisterModule("/m3", m3.Routes)
	s.RegisterModule("/m4", m4.Routes)

	// start the server
	s.Start()
}
