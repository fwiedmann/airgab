package rsync

import (
	"os"
	"os/exec"
)

// Rsync for client
type Rsync struct {
	user        string
	host        string
	source      string
	destination string
	options     string // tbd. string array
}

// checkKey  = For ssh-connection a private-key is needed.
func checkKey() error {
	_, err := os.Open("/home/fwiedmann/.ssh/id_rsa")
	if err != nil {
		return err
	}
	return nil
}

// CheckConnection = A rsync dry-run will test the ssh connection and backup functionality
func (r *Rsync) CheckConnection() {
	cmd := exec.Command("rsync", "--dry-run", r.user+"@"+r.host+":"+r.source, r.destination)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// RunSync  = Pass rsync options as one string.
func (r *Rsync) RunSync() {
	cmd := exec.Command("rsync", r.options, r.user+"@"+r.host+":"+r.source, r.destination)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// InitRsync = Required information for the rsync client
func InitRsync(user, host, source, destination, options string) *Rsync {
	if err := checkKey(); err != nil {
		panic(err)
	}
	return &Rsync{user, host, source, destination, options}
}
