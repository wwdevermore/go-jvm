package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	stack.PushSlot(s1)
	stack.PushSlot(s2)
}
