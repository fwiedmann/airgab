package opts

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug       = kingpin.Flag("debug", "Enable debug mode.").Bool()
	host        = kingpin.Flag("host", "Host from where the data will be backuped (user@host).").Default().String()
	source      = kingpin.Flag("source", "Absolute path to directory from where the data will be backuped.").Required().String()
	destination = kingpin.Flag("destination", "Absolute path to directory whre the data will be stored.").Required().String()
)

// Opts from user input
type Opts struct {
	Debug               bool
	Source, Destination string
}

func formatSource(host, source *string) string {
	if *host == "" {
		return *source
	}
	return fmt.Sprintf("%v:%v", *host, *source)
}

// New returns opts struct for use in main
func New() *Opts {
	kingpin.Parse()
	return &Opts{*debug, formatSource(host, source), *destination}
}
