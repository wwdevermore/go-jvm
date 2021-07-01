package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
