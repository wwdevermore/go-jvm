package references

import (
	"fmt"
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
	"go-jvm/ch05/rtda/heap"
)

//调用的方法取决于实例的实际类型
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		if methodRef.Name() == "println" {
			stack := frame.OperandStack()
			switch methodRef.Descriptor() {
			case "(Z)V":
				fmt.Printf("%v\n", stack.PopInt() != 0)
			case "(C)V":
				fmt.Printf("%c\n", stack.PopInt())
			case "(B)V":
				fmt.Printf("%v\n", stack.PopInt())
			case "(S)V":
				fmt.Printf("%v\n", stack.PopInt())
			case "(I)V":
				fmt.Printf("%v\n", stack.PopInt())
			case "(J)V":
				fmt.Printf("%v\n", stack.PopLong())
			case "(F)V":
				fmt.Printf("%v\n", stack.PopFloat())
			case "(D)V":
				fmt.Printf("%v\n", stack.PopDouble())
			default:
				panic("println: " + methodRef.Descriptor())
			}
			stack.PopRef()
			return
		}
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAceessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
