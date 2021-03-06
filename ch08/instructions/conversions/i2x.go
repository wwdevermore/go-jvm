package conversions

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type I2L struct {
	base.NoOperandsInstruction
}

type I2D struct {
	base.NoOperandsInstruction
}

type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopInt()
	d := float64(l)
	stack.PushDouble(d)
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}
