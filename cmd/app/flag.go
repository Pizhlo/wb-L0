package main

import (
	"flag"
)

type Config struct {
	FlagDBAddress string
}

func ParseFlags() Config {
	var conf Config

	flag.StringVar(&conf.FlagDBAddress, "d", "", "database address")

	flag.Parse()

	return conf
}
