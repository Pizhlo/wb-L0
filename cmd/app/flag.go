package main

import (
	"flag"
)

type Config struct {
	DBAddress string
	MigratePath   string
}

func ParseFlags() Config {
	var conf Config

	flag.StringVar(&conf.DBAddress, "d", "", "database address")
	flag.StringVar(&conf.MigratePath, "m", "", "migrate path")

	flag.Parse()

	return conf
}
