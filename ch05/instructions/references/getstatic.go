package references

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"go-jvm/ch05/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (receiver *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(receiver.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.Descriptor()
	stack := frame.OperandStack()
	slotId := field.SlotId()
	slots := class.StaticVars()
	switch descriptor[0] {
	case 'Z','B','C','S','I': stack.PushInt(slots.GetInt(slotId))
	case 'F': stack.PushFloat(slots.GetFloat(slotId))
	case 'J': stack.PushLong(slots.GetLong(slotId))
	case 'D': stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[': stack.PushRef(slots.GetRef(slotId))
	}
}
