package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
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

func main() {
	masterServerIP = flag.String("master", "127.0.0.1:43210", "master server ip address (127.0.0.1:43210)")
	chunkServerIP = flag.String("chunk", "127.0.0.1:43300", "chunk server ip address (127.0.0.1:43300)")
	flag.Parse()

	startChunkServer(*chunkServerIP)
}

func startChunkServer(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/add", Add)
	r.HandleFunc("/get", Get)
	http.Handle("/", r)

	glog.Infof("chunk listen at %v", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		glog.Fatal(err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	log.Println("add")

	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File["uploadfile"]
	for _, fileHeader := range files {

		randkey := uint64(RangKeyGenerator)
		RangKeyGenerator++

		fileInfo, _ := ChunkFile.Stat()

		offset := fileInfo.Size()

		file, err := fileHeader.Open()
		if err != nil {
			log.Println(err)
			continue
		}

		written, err := io.Copy(ChunkFile, file)
		if err != nil {
			log.Println(err)
			continue
		}

		if written != fileHeader.Size {
			log.Println(written, fileHeader.Size)
			continue
		}

		m[randkey] = TLocation{
			Offset: uint32(offset),
			Size:   uint32(written),
		}

		log.Println(randkey, offset, written)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Println("get")

	r.ParseForm()

	log.Println(r.Form, r.PostForm)

	randkeystr := r.Form.Get("key")
	randkey, _ := strconv.Atoi(randkeystr)

	location, ok := m[uint64(randkey)]
	if !ok {
		log.Println("invalid randkey = ", randkey)
		return
	}

	b := make([]byte, location.Size)
	n, err := ChunkFile.ReadAt(b, int64(location.Offset))
	if err != nil {
		log.Println(err)
		return
	}

	if uint32(n) != location.Size {
		log.Println(n, location.Size)
		return
	}

	buffer := bytes.NewBuffer(b)

	io.Copy(w, buffer)
}
