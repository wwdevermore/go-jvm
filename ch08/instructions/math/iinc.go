package math

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	vars := frame.LocalVars()
	result := vars.GetInt(self.Index) + self.Const
	vars.SetInt(self.Index, result)
}
