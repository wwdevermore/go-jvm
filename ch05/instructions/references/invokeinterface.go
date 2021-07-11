package references

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"go-jvm/ch05/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	//count uint8
	//zero uint8
}

func (receiver *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	receiver.index = uint(reader.ReadInt16())
	reader.ReadInt8() //count
	reader.ReadInt8() //must be 0
}

func (receiver *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(receiver.index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(method.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplementsOf(method.Class()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
