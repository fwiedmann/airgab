package opts

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug       = kingpin.Flag("debug", "Enable debug mode.").Bool()
	user        = kingpin.Flag("user", "Host from which the data will be backuped").PlaceHolder("USER").Required().String()
	host        = kingpin.Flag("host", "Host from which the data will be backuped").PlaceHolder("HOSTNAME").Required().String()
	port        = kingpin.Flag("port", "Port for ssh connection").PlaceHolder("default: 22").Default("22").Int()
	source      = kingpin.Flag("source", "Absolute path to directory from where the data will be backuped.").PlaceHolder("/source/directory").Required().String()
	destination = kingpin.Flag("destination", "Absolute path to directory whre the data will be stored.").PlaceHolder("/destination/directory").Required().String()
)

// Opts from user input
type Opts struct {
	Debug                           bool
	User, Host, Source, Destination string
	Port                            int
}

// New returns opts struct for use in main
func New() *Opts {
	kingpin.Parse()
	return &Opts{*debug, *user, *host, *source, *destination, *port}
}
