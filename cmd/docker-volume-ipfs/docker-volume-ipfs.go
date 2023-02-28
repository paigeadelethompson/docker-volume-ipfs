package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/paigeadelethompson/docker-volume-ipfs/driver"
	"github.com/paigeadelethompson/docker-volume-ipfs/kubo"
	"github.com/paigeadelethompson/docker-volume-ipfs/fuse"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	k := new(kubo.KuboController)
	f := new(fuse.FUSEController)
	d := driver.New("/mnt/docker_ipfs_mounts", f, k)
	h := volume.NewHandler(&d)
	
	go func() {
		if err := h.ServeUnix("root", os.Getgid()); err != nil {
			log.Fatal(err)
		}
	}()
}
