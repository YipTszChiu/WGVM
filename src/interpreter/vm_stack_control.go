package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/common/op"
)

// 调用帧
type controlFrame struct {
	opcode byte
	bt     common.FuncType      // block type
	instrs []common.Instruction // expr
	bp     int                  // base pointer (operand stack)
	pc     int                  // program counter
}

// 构造函数
func newControlFrame(opcode byte,
	bt common.FuncType, instrs []common.Instruction, bp int) *controlFrame {

	return &controlFrame{
		opcode: opcode,
		bt:     bt,
		instrs: instrs,
		bp:     bp,
		pc:     0,
	}
}

// 调用栈
type controlStack struct {
	frames []*controlFrame
}

func (cs *controlStack) pushControlFrame(cf *controlFrame) {
	cs.frames = append(cs.frames, cf)
}
func (cs *controlStack) popControlFrame() *controlFrame {
	cf := cs.frames[len(cs.frames)-1]
	cs.frames = cs.frames[:len(cs.frames)-1]
	return cf
}

// 获取控制栈深度
func (cs *controlStack) controlDepth() int {
	return len(cs.frames)
}

// 获取栈顶控制帧
func (cs *controlStack) topControlFrame() *controlFrame {
	return cs.frames[len(cs.frames)-1]
}

// 最靠近栈顶的调用帧
func (cs *controlStack) topCallFrame() (*controlFrame, int) {
	for n := len(cs.frames) - 1; n >= 0; n-- {
		if cf := cs.frames[n]; cf.opcode == op.Call {
			return cf, len(cs.frames) - 1 - n
		}
	}
	return nil, -1
}
