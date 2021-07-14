package references

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (receiver *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	receiver.index = reader.ReadUint16()
	receiver.dimensions = reader.ReadUint8()
}

func (receiver *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(receiver.index)).(*heap.ClassRef)
	arrClass := classRef.ResolveClass()
	counts := popAndCheckCounts(stack, int(receiver.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}
