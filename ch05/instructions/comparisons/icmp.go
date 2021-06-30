package comparisons

import (
	"go-jvm/ch04/rtda"
	"go-jvm/ch05/instructions/base"
)

type ICMP struct {
	base.NoOperandsInstruction
}

func (self ICMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
