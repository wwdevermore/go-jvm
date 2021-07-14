package loads

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

type AALOAD struct {
	base.NoOperandsInstruction
}
type BALOAD struct {
	base.NoOperandsInstruction
}
type CALOAD struct {
	base.NoOperandsInstruction
}
type DALOAD struct {
	base.NoOperandsInstruction
}
type FALOAD struct {
	base.NoOperandsInstruction
}
type IALOAD struct {
	base.NoOperandsInstruction
}
type LALOAD struct {
	base.NoOperandsInstruction
}
type SALOAD struct {
	base.NoOperandsInstruction
}

func (receiver *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Refs()[arrIndex]
	stack.PushRef(val)
}
func (receiver *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Bytes()[arrIndex]
	stack.PushInt(int32(val))
}
func (receiver *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Chars()[arrIndex]
	stack.PushInt(int32(val))
}
func (receiver *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Doubles()[arrIndex]
	stack.PushDouble(val)
}
func (receiver *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Floats()[arrIndex]
	stack.PushFloat(val)
}
func (receiver *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Ints()[arrIndex]
	stack.PushInt(val)
}
func (receiver *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Longs()[arrIndex]
	stack.PushLong(val)
}
func (receiver *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrIndex := stack.PopInt()
	arrRef := stack.PopRef()
	checkNilObject(arrRef)
	checkIndex(arrIndex, arrRef)
	val := arrRef.Shorts()[arrIndex]
	stack.PushInt(int32(val))
}

func checkIndex(arrIndex int32, arrRef *heap.Object) {
	if arrIndex >= arrRef.ArrayLength() || arrIndex < 0 {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}

func checkNilObject(arrRef *heap.Object) {
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
}
