package common

import "WGVM/src/common/op"

type Expr = []Instruction

type Instruction struct {
	Opcode byte
	Args   interface{}
}

// block和loop指令立即数
type BlockArgs struct {
	BT     BlockType
	Instrs []Instruction
}

// if指令立即数
type IfArgs struct {
	BT      BlockType
	Instrs1 []Instruction
	Instrs2 []Instruction
}

// br_table指令
type BrTableArgs struct {
	Labels  []LabelIdx
	Default LabelIdx
}

// 内存指令立即数
type MemArg struct {
	// 对齐提示
	Align uint32
	// 内存偏移量
	Offset uint32
}

func (instr Instruction) GetOpname() string {
	return op.Opnames[instr.Opcode]
}
func (instr Instruction) String() string {
	return op.Opnames[instr.Opcode]
}
