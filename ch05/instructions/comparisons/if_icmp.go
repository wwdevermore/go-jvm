package comparisons

//比较操作数栈前两个数的大小
import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

type IF_ICMPNE struct {
	base.BranchInstruction
}
type IF_ICMPLT struct {
	base.BranchInstruction
}
type IF_ICMPLE struct {
	base.BranchInstruction
}
type IF_ICMPGT struct {
	base.BranchInstruction
}
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
