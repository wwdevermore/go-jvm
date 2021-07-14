package stores

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

type AASTORE struct {
	base.NoOperandsInstruction
}
type BASTORE struct {
	base.NoOperandsInstruction
}
type CASTORE struct {
	base.NoOperandsInstruction
}
type DASTORE struct {
	base.NoOperandsInstruction
}
type FASTORE struct {
	base.NoOperandsInstruction
}
type SASTORE struct {
	base.NoOperandsInstruction
}
type IASTORE struct {
	base.NoOperandsInstruction
}
type LASTORE struct {
	base.NoOperandsInstruction
}

func (receiver *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopRef()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Refs()[assignIndex] = valToAssign
}

func (receiver *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopInt()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Bytes()[assignIndex] = int8(valToAssign)
}

func (receiver *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopInt()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Chars()[assignIndex] = uint16(valToAssign)
}

func (receiver *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopDouble()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Doubles()[assignIndex] = valToAssign
}

func (receiver *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopFloat()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Floats()[assignIndex] = valToAssign
}

func (receiver *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopInt()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Shorts()[assignIndex] = int16(valToAssign)
}

func (receiver *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopInt()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Ints()[assignIndex] = valToAssign
}

func (receiver *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	valToAssign := stack.PopLong()
	assignIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkIndex(assignIndex, arrRef)
	checkNotNil(arrRef)
	arrRef.Longs()[assignIndex] = valToAssign
}

func checkIndex(assignIndex int32, arrRef *heap.Object) {
	if assignIndex < 0 || assignIndex >= arrRef.ArrayLength() {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(arrRef *heap.Object) {
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
}
