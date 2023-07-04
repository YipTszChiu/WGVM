package interpreter

import "math"

// 操作数栈
type operandStack struct {
	slots []uint64
}

func (s *operandStack) PushU64(val uint64) {
	s.slots = append(s.slots, val)
}
func (s *operandStack) PopU64() uint64 {
	val := s.slots[len(s.slots)-1]
	s.slots = s.slots[:len(s.slots)-1]
	return val
}

func (s *operandStack) PushS64(val int64) {
	s.PushU64(uint64(val))
}
func (s *operandStack) PopS64() int64 {
	return int64(s.PopU64())
}

func (s *operandStack) PushU32(val uint32) {
	s.PushU64(uint64(val))
}
func (s *operandStack) PopU32() uint32 {
	return uint32(s.PopU64())
}

func (s *operandStack) PushS32(val int32) {
	s.PushU32(uint32(val))
}
func (s *operandStack) PopS32() int32 {
	return int32(s.PopU32())
}

func (s *operandStack) PushF32(val float32) {
	s.PushU32(math.Float32bits(val))
}
func (s *operandStack) PopF32() float32 {
	return math.Float32frombits(s.PopU32())
}

func (s *operandStack) PushF64(val float64) {
	s.PushU64(math.Float64bits(val))
}
func (s *operandStack) PopF64() float64 {
	return math.Float64frombits(s.PopU64())
}

func (s *operandStack) PushBool(val bool) {
	if val {
		s.PushU64(1)
	} else {
		s.PushU64(0)
	}
}
func (s *operandStack) PopBool() bool {
	return s.PopU64() != 0
}
