package base

import "go-jvm/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index int8
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = int8(reader.ReadInt8())
}

type Index16Instruction struct {
	Index int16
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = int16(reader.ReadInt16())
}
