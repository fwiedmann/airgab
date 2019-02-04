package main

import "github.com/fwiedmann/airgab/pkg/rsync"

func main() {
	r := rsync.InitRsync("pi", "192.168.2.233", "22", "/home/pi/test.txt", ".")
	r.CheckConnection()
}
