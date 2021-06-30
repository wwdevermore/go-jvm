package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IFEQ struct {
	base.NoOperandsInstruction
}

type IFNE struct {
	base.NoOperandsInstruction
}

type IFLT struct {
	base.NoOperandsInstruction
}

type IFLE struct {
	base.NoOperandsInstruction
}

type IFGT struct {
	base.NoOperandsInstruction
}

type IFGE struct {
	base.NoOperandsInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
