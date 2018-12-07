package master

type TChunkServer struct {
	ID   uint64
	Addr string //"127.0.0.1:1234"
}

type TMetaData struct {
}

var GMeta TMetaData

func LoadMetaFromFile() {
}

func SaveMetaToFile() {
}
