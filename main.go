package main

import (
	"github.com/fwiedmann/airgab/pkg/opts"
	"github.com/fwiedmann/airgab/pkg/rsync"
)

func main() {
	opts := opts.New()
	r := rsync.InitRsync(opts.User, opts.Host, opts.Source, opts.Destination, opts.Options)
	r.CheckConnection()
	r.RunSync()
}
