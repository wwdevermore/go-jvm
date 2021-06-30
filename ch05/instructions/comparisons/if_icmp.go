package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IF_ICMPEQ struct {
	base.NoOperandsInstruction
}

type IF_ICMPNE struct {
	base.NoOperandsInstruction
}
type IF_ICMPLT struct {
	base.NoOperandsInstruction
}
type IF_ICMPLE struct {
	base.NoOperandsInstruction
}
type IF_ICMPGT struct {
	base.NoOperandsInstruction
}
type IF_ICMPGE struct {
	base.NoOperandsInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		frame.Branch(frame, self.Offset)
	}
}
