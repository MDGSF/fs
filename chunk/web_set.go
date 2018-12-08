package chunk

import (
	"io"
	"log"
	"net/http"
)

func Set(w http.ResponseWriter, r *http.Request) {
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
