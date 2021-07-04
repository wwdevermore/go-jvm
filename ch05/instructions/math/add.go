package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IADD struct {
	base.NoOperandsInstruction
}

type LADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value2 := stack.PopInt()
	value1 := stack.PopInt()
	stack.PushInt(value2 + value1)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value2 := stack.PopLong()
	value1 := stack.PopLong()
	stack.PushLong(value2 + value1)
}
