package chunk

import (
	"flag"
	"os"
	"os/signal"

	"github.com/golang/glog"
)

type TLocation struct {
	Offset uint32
	Size   uint32
}

var ChunkFile *os.File

var m map[uint64]TLocation

var RangKeyGenerator uint64

var masterServerIP *string
var chunkServerIP *string
var volumeDir *string
var maxVolumeNum *int

// Start start chunk server
func Start() {
	masterServerIP = flag.String("master", "127.0.0.1:43210", "master server ip address (127.0.0.1:43210)")
	chunkServerIP = flag.String("chunk", "127.0.0.1:43300", "chunk server ip address (127.0.0.1:43300)")
	volumeDir = flag.String("dir", "/tmp", "directory to store file (/tmp)")
	maxVolumeNum = flag.Int("max", 3, "maximum numbers of volume")
	flag.Parse()

	initVolumeDir()

	startWebService()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	os.Exit(0)
}

func initVolumeDir() {
	fileInfo, err := os.Stat(*volumeDir)
	if err != nil {
		if err := os.MkdirAll(*volumeDir, 0775); err != nil {
			glog.Fatalf("mkdir %v failed, err = %v", *volumeDir, err)
		}
		return
	}

	if !fileInfo.IsDir() {
		glog.Fatalf("%v is not directory.", *volumeDir)
	}
}
