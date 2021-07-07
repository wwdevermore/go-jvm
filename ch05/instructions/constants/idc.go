package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }

func (receiver *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, receiver.Index)
}

func (receiver *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, receiver.Index)
}

func (receiver *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(receiver.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))

	// case string:
	// case *heap.ClassRef:
	default:
		panic("todo: ldc !")
	}
}
