package comparisons

//比较操作数栈第一个值和零的大小
import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}

type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		//fmt.Printf("val: %d greater than 0, offset %d\n", val, self.Offset)
		base.Branch(frame, self.Offset)
	}
}
