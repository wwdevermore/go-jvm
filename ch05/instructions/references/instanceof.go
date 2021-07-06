package references

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"go-jvm/ch05/rtda/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (receiver *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	classRef := frame.Method().Class().ConstantPool().GetConstant(receiver.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
