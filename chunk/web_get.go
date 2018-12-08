package chunk

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
)

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
