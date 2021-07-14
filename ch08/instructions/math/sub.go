package math

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type LSUB struct {
	base.NoOperandsInstruction
}
type ISUB struct {
	base.NoOperandsInstruction
}

func (receiver *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}

func (receiver *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}
