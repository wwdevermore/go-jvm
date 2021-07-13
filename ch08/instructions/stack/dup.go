package stack

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type DUP struct {
	base.NoOperandsInstruction
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

type DUP2 struct {
	base.NoOperandsInstruction
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot3 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot3 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot3)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot4 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot3)
	stack.PushSlot(slot4)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
