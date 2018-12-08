package chunk

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*
volume_id.idx
file id(64) | offset(32) | size(32)
*/

const (
	FileIDBytes     = 8
	FileOffsetBytes = 4
	FileSizeBytes   = 4
	SectionBytes    = FileIDBytes + FileOffsetBytes + FileSizeBytes
)

type Section struct {
	FileID     uint64
	FileOffset uint32
	FileSize   uint32
}

func LoadIdxFile(name string) {
}

func CreateVolumeIdx(name string) (*os.File, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func WriteToVolumIdx(f *os.File, fid uint64, offset uint32, size uint32) error {
	sec := make([]byte, SectionBytes)
	binary.BigEndian.PutUint64(sec, fid)
	binary.BigEndian.PutUint32(sec[FileIDBytes:], offset)
	binary.BigEndian.PutUint32(sec[FileIDBytes+FileOffsetBytes:], size)
	written, err := f.Write(sec)
	if err != nil {
		return err
	}
	if written != SectionBytes {
		return fmt.Errorf("written = %v", written)
	}
	return nil
}
