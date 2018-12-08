package chunk

import "fmt"

type Volume struct {
	ID              uint32
	mapFile2Section map[uint64]*Section
}

func NewVolume(ID uint32) *Volume {
	v := &Volume{}
	v.mapFile2Section = make(map[uint64]*Section)
	return v
}

func (v *Volume) LoadFromFile() {
}

func (v *Volume) IdxName() string {
	return fmt.Sprintf("volume_%v.idx", v.ID)
}

func (v *Volume) HeaderName() string {
	return fmt.Sprintf("volume_%v.header", v.ID)
}

func (v *Volume) DatName() string {
	return fmt.Sprintf("volume_%v.dat", v.ID)
}
