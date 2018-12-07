package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test16BE(t *testing.T) {
	b := Uint16toBytesBE(0x1234)
	num := BytesToUint16BE(b)
	assert.Equal(t, uint16(0x1234), num, "they should be equal")
}

func Test32BE(t *testing.T) {
	b := Uint32toBytesBE(0x1234)
	num := BytesToUint32BE(b)
	assert.Equal(t, uint32(0x1234), num, "they should be equal")
}

func Test64BE(t *testing.T) {
	b := Uint64toBytesBE(0x1234)
	num := BytesToUint64BE(b)
	assert.Equal(t, uint64(0x1234), num, "they should be equal")
}

func Test16LE(t *testing.T) {
	b := Uint16toBytesLE(0x1234)
	num := BytesToUint16LE(b)
	assert.Equal(t, uint16(0x1234), num, "they should be equal")
}

func Test32LE(t *testing.T) {
	b := Uint32toBytesLE(0x1234)
	num := BytesToUint32LE(b)
	assert.Equal(t, uint32(0x1234), num, "they should be equal")
}

func Test64LE(t *testing.T) {
	b := Uint64toBytesLE(0x1234)
	num := BytesToUint64LE(b)
	assert.Equal(t, uint64(0x1234), num, "they should be equal")
}
