package references

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (receiver *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	arrRef := frame.OperandStack().PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	frame.OperandStack().PushInt(arrLen)
}
