package comparisons

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type FCMPL struct {
	base.NoOperandsInstruction
}

type FCMPG struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
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

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}
