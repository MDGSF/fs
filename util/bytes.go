package util

import "encoding/binary"

func BytesToUint16BE(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func BytesToUint32BE(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func BytesToUint64BE(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func Uint16toBytesBE(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func Uint32toBytesBE(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

func Uint64toBytesBE(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func BytesToUint16LE(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func BytesToUint32LE(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func BytesToUint64LE(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Uint16toBytesLE(v uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}

func Uint32toBytesLE(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func Uint64toBytesLE(v uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
}
