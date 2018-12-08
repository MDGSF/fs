package master

import (
	"flag"
	"net/http"

	"github.com/MDGSF/fs/numgener"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

var masterServerIP *string

var FileIDGener *numgener.TNumGenerator

func Start() {
	masterServerIP = flag.String("master", "127.0.0.1:43210", "master server ip address (127.0.0.1:43210)")
	flag.Parse()

	FileIDGener = numgener.NewNumGenerator()

	r := mux.NewRouter()
	r.HandleFunc("/ws/chunk", chunkHandler)
	r.HandleFunc("/dir/assign", clientHandler).Methods("GET")
	http.Handle("/", r)

	glog.Infof("master listen at %v", *masterServerIP)
	if err := http.ListenAndServe(*masterServerIP, nil); err != nil {
		glog.Fatal(err)
	}
}
