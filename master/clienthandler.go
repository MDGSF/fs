package master

import (
	"encoding/json"
	"net/http"
)

func clientHandler(w http.ResponseWriter, r *http.Request) {
	newFileID := FileIDGener.Num()
	m := make(map[string]interface{})
	m["fid"] = newFileID
	b, _ := json.Marshal(m)
	w.Write(b)
}
