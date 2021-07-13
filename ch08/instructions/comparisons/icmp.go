package comparisons

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
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
