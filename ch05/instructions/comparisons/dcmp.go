package comparisons

import (
	"go-jvm/ch04/rtda"
	"go-jvm/ch05/instructions/base"
)

type DCMPL struct {
	base.NoOperandsInstruction
}

type DCMPG struct {
	base.NoOperandsInstruction
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}
