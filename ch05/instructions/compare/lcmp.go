package compare

import (
	"go-jvm/ch04/rtda"
	"go-jvm/ch05/instructions/base"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
