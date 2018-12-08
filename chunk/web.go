package chunk

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func startWebService() {
	r := mux.NewRouter()
	r.HandleFunc("/set", Set)
	r.HandleFunc("/get", Get)

	srv := http.Server{
		Addr:         *chunkServerIP,
		WriteTimeout: time.Second * 15,
		Handler:      r,
	}

	go func() {
		glog.Infof("chunk listen at %v", *chunkServerIP)
		if err := srv.ListenAndServe(); err != nil {
			glog.Fatal(err)
		}
	}()
}
