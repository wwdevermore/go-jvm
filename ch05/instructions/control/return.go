package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type RETURN struct { base.NoOperandsInstruction }
type ARETURN struct { base.NoOperandsInstruction }
type DRETURN struct { base.NoOperandsInstruction }
type FRETURN struct { base.NoOperandsInstruction }
type IRETURN struct { base.NoOperandsInstruction }
type LRETURN struct { base.NoOperandsInstruction }

func (receiver *RETURN) Execute(frame *rtda.Frame)  {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}