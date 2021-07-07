package references

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (receiver *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
