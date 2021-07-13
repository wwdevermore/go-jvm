package control

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
