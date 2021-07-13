package references

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

//SuperClass.interset(),SubClass.interset(), invoke special会调用的方法取决于引用的类型
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (receiver *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	methodRef := currentClass.ConstantPool().GetConstant(receiver.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	class := methodRef.ResolveClass()
	if method.Name() == "<init>" && method.Class() != class {
		panic("java.lang.NoSuchMethodError")
	}
	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(method.ArgSlotCount() - 1)
	if method.IsProtected() &&
		method.Class().IsSuperClassOf(currentClass) &&
		method.Class().GetPackageName() != currentClass.GetPackageName() &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegaAccessError")
	}
	methodToBeInvoked := method
	if currentClass.IsSupper() &&
		class.IsSuperClassOf(currentClass) &&
		method.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
