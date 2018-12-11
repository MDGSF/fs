package master

type TChunkServer struct {
	ID   uint64
	Addr string //"127.0.0.1:1234"
}

type TSmallFile struct {
	ID          uint64
	ChunkServer []string
}

type TBigFile struct {
	ID    uint64
	Parts []TSmallFile
}

func LoadMetaFromFile() {
}

func SaveMetaToFile() {
}
