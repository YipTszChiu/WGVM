package interpreter

import (
	"WGVM/src/common"
	"math"
)

func wrapU64(vt common.ValType, val uint64) interface{} {
	switch vt {
	case common.ValTypeI32:
		return int32(val)
	case common.ValTypeI64:
		return int64(val)
	case common.ValTypeF32:
		return math.Float32frombits(uint32(val))
	case common.ValTypeF64:
		return math.Float64frombits(val)
	default:
		panic("unreachable") // TODO
	}
}

func unwrapU64(vt common.ValType, val interface{}) uint64 {
	switch vt {
	case common.ValTypeI32:
		return uint64(val.(int32))
	case common.ValTypeI64:
		return uint64(val.(int64))
	case common.ValTypeF32:
		return uint64(math.Float32bits(val.(float32)))
	case common.ValTypeF64:
		return math.Float64bits(val.(float64))
	default:
		panic("unreachable") // TODO
	}
}
