package rsync

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fwiedmann/airgab/pkg/opts"
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
func (r *Rsync) CheckKey() error {
	_, err := os.Open(r.sshkey)
	if err != nil {
		return err
	}
	return nil
}

// CheckConnection  A rsync dry-run will test the ssh connection and backup functionality
func (r *Rsync) CheckConnection() {
	cmd := exec.Command("rsync", "--dry-run", r.user+"@"+r.host+":"+r.source, r.destination)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// RunSync  Pass rsync options as one string.
func (r *Rsync) RunSync() {
	cmd := exec.Command("rsync", r.options, r.user+"@"+r.host+":"+r.source, r.destination)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
