package opts

import (
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug       = kingpin.Flag("debug", "Enable debug mode.").Bool()
	user        = kingpin.Flag("user", "Host from which the data will be backuped").PlaceHolder("USER").Required().String()
	host        = kingpin.Flag("host", "Host from which the data will be backuped").PlaceHolder("HOSTNAME").Required().String()
	source      = kingpin.Flag("source", "Absolute path to directory from where the data will be backuped.").PlaceHolder("/source/directory").Required().String()
	destination = kingpin.Flag("destination", "Absolute path to directory whre the data will be stored.").PlaceHolder("/destination/directory").Required().String()
	options     = kingpin.Flag("options", "rsync options").PlaceHolder("-a").Default("-a").String()
	interval    = kingpin.Flag("interval", "interval of airgab run").Required().String()
	sshkey      = kingpin.Flag("private-key", "absolute path to private ssh key").Default("/home/user/.ssh/id_rsa").String()
)

// Opts from user input
type Opts struct {
	Debug                                            bool
	User, Host, Source, Destination, Options, Sshkey string
	Interval                                         time.Duration
}

// New returns opts struct for use in main
func New() *Opts {
	kingpin.Parse()

	parsedInterval, err := time.ParseDuration(*interval)
	if err != nil {
		panic(err)
	}

	return &Opts{*debug, *user, *host, *source, *destination, *options, *sshkey, parsedInterval}
}
