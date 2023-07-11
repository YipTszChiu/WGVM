package interpreter

import "WGVM/src/common"

type memory struct {
	_type common.MemType
	data  []byte
}

func newMemory(mt common.MemType) *memory {
	return &memory{
		_type: mt,
		data:  make([]byte, mt.Min*common.PageSize),
	}
}

func (mem *memory) Type() common.MemType {
	return mem._type
}

func (mem *memory) Size() uint32 {
	return uint32(len(mem.data) / common.PageSize)
}
func (mem *memory) Grow(n uint32) uint32 {
	oldSize := mem.Size()
	if n == 0 {
		return oldSize
	}

	maxPageCount := uint32(common.MaxPageCount)
	if max := mem._type.Max; max > 0 {
		maxPageCount = max
	}
	if oldSize+n > maxPageCount {
		return 0xFFFFFFFF // -1
	}

	newData := make([]byte, (oldSize+n)*common.PageSize)
	copy(newData, mem.data)
	mem.data = newData
	return oldSize
}

func (mem *memory) Read(offset uint64, buf []byte) {
	mem.checkOffset(offset, len(buf))
	copy(buf, mem.data[offset:])
}
func (mem *memory) Write(offset uint64, data []byte) {
	mem.checkOffset(offset, len(data))
	copy(mem.data[offset:], data)
}

func (mem *memory) checkOffset(offset uint64, length int) {
	if int64(len(mem.data)-length) < int64(offset) {
		panic(ErrMemOutOfBounds)
	}
}
