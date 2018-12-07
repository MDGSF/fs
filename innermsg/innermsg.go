package innermsg

// HeartBeat chunk server send to master server.
type HeartBeat struct {
	ID      uint64        `msgpack:"id"`
	Volumes []interface{} `msgpack:"volumes"`
}
