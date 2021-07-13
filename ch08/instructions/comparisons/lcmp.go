package comparisons

import (
	"fmt"
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v2 > v1 {
		fmt.Printf("%d > %d, push 1 onto stack\n", v2, v1)
		stack.PushInt(1)
	} else if v2 < v1 {
		fmt.Printf("%d < %d, push -1 onto stack\n", v2, v1)
		stack.PushInt(-1)
	} else {
		fmt.Printf("%d = %d, push 0 onto stack\n", v2, v1)
		stack.PushInt(0)
	}
}
