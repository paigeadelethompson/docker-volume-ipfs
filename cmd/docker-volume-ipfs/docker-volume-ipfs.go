package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/paigeadelethompson/docker-volume-ipfs/driver"
)

const ipfsID = "_ipfs"

var (
	defaultDir     = filepath.Join(volume.DefaultDockerRootDirectory, ipfsID)
	ipfsMountPoint = flag.String("mount", "/ipfs", "ipfs mount point")
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	flag.Parse()

	_, err := os.Lstat(*ipfsMountPoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n%s does not exists, can't start..\n Please use ipfs command line to mount it\n", err, *ipfsMountPoint)
		os.Exit(1)
	}

	d := driver.New(*ipfsMountPoint)
	h := volume.NewHandler(d)

	go func() {
		if err := h.ServeUnix("root", os.Getegid()); err != nil {
			fmt.Println(err)
		}
	}()

	setupKubo()
}
