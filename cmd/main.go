package main

import (
	"flag"
	"playground/config"
)

var configDir = flag.String("confdir", ".", "Config file location")
var configFile = flag.String("confname", "config", "Config file name (w/o extension)")

func main() {
	conf, err := config.GetConfig(*configDir, *configFile)
	panicErr(err)

	conf.Token.Token
}

func panicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
