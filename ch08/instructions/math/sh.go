package math

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type ISHL struct { //算数左移
	base.NoOperandsInstruction
}

type ISHR struct { //算数右移
	base.NoOperandsInstruction
}

type IUSHR struct { //逻辑左移
	base.NoOperandsInstruction
}

type LSHL struct {
	base.NoOperandsInstruction
}

type LSHR struct {
	base.NoOperandsInstruction
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := int32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := int32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x1f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
