package opts

import (
	"time"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	user        = kingpin.Flag("user", "Host from which the data will be backuped").PlaceHolder("USER").Required().String()
	host        = kingpin.Flag("host", "Host from which the data will be backuped").PlaceHolder("HOSTNAME").Required().String()
	source      = kingpin.Flag("source", "Absolute path to directory from where the data will be backuped.").PlaceHolder("/source/directory").Required().String()
	destination = kingpin.Flag("destination", "").PlaceHolder("/destination/directory").Default("/home/pilot/backup").String()
	options     = kingpin.Flag("options", "rsync options").PlaceHolder("-a").Default("-a").String()
	interval    = kingpin.Flag("interval", "interval of airgab run with time unit(s,m,h) ").PlaceHolder("10m").Required().String()
	sshkey      = kingpin.Flag("private-key", "absolute path to private ssh key").Default("/home/pilot/.ssh/id_rsa").String()
)

// Opts from user input
type Opts struct {
	User, Host, Source, Destination, Options, Sshkey string
	Interval                                         time.Duration
}

// New returns opts struct for use in main
func New() *Opts {
	kingpin.Parse()

	parsedInterval, err := time.ParseDuration(*interval)
	if err != nil {
		log.Panic("Please check the format of the interval option.")
	}

	return &Opts{*user, *host, *source, *destination, *options, *sshkey, parsedInterval}
}
