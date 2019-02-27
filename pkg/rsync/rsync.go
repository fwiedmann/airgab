package rsync

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fwiedmann/airgab/pkg/opts"
	log "github.com/sirupsen/logrus"
)

// Rsync for client
type Rsync struct {
	user, host, source, destination, options, sshkey string
}

// InitRsync  Required information for the rsync client
func InitRsync(opts opts.Opts) *Rsync {
	return &Rsync{opts.User, opts.Host, opts.Source, opts.Destination, opts.Options, opts.Sshkey}
}

// GetBackupSize returns the current backupsize in megabyte
func (r *Rsync) GetBackupSize() float64 {
	var files []int64
	var size int64

	err := filepath.Walk(r.destination, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Size())
			return nil
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		size += file
	}
	return float64(size / 1000000)
}

// CheckKey checks if private key is existing
func (r *Rsync) CheckKey() {
	log.Infof("Check if private key is existing under: %s.", r.sshkey)
	_, err := os.Open(r.sshkey)
	if err != nil {
		log.Panic("Cant find or open private ssh-key.")
	}
}

// CheckConnection  A rsync dry-run will test the ssh connection and backup functionality
func (r *Rsync) CheckConnection() {
	log.Info("Check connection to source host.")
	cmd := exec.Command("rsync", "--dry-run", "-e ssh '-o StrictHostKeyChecking=no'", r.user+"@"+r.host+":"+r.source, r.destination)
	if err := cmd.Run(); err != nil {
		log.Panic("Failed to connect to source host. Please check your connection settings and ssh-key.")
	}
}

// RunSync  Pass rsync options as one string.
func (r *Rsync) RunSync() float64 {
	cmd := exec.Command("rsync", "-e ssh '-o StrictHostKeyChecking=no'", r.options, r.user+"@"+r.host+":"+r.source, r.destination)
	timeBeforeSync := time.Now()

	if err := cmd.Run(); err != nil {
		log.Panicf("Failed to backup %s. Please check your connection settings and ssh-key.", r.source)
		panic(err)
	}

	timeAfterSync := time.Now()
	durationInMinutes := 166.667 * (float64(timeAfterSync.Sub(timeBeforeSync)) / float64(time.Second))
	return durationInMinutes
}
