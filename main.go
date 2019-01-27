package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug       = kingpin.Flag("debug", "Enable debug mode.").Bool()
	host        = kingpin.Flag("host", "Host from where the data will be backuped (user@host).").Default("").String()
	source      = kingpin.Flag("source", "Absolute path to directory from where the data will be backuped.").Required().String()
	destination = kingpin.Flag("destination", "Absolute path to directory whre the data will be stored.").Required().String()
)

func main() {
	kingpin.Parse()
}
