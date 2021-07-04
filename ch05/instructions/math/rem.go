package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}

type FREM struct {
	base.NoOperandsInstruction
}

type IREM struct {
	base.NoOperandsInstruction
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException by zero \n\r")
	}
	result := val1 % val2
	stack.PushInt(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException by zero \n\r")
	}
	result := val1 % val2
	stack.PushLong(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := float64(stack.PopFloat())
	val2 := float64(stack.PopFloat())
	result := math.Mod(val1, val2)
	stack.PushFloat(float32(result))
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}
