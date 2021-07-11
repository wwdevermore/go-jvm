package rtda

import "go-jvm/ch05/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
	method       *heap.Method
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (receiver *Frame) RevertNextPC() {
	receiver.nextPC = receiver.thread.pc
}

func (receiver *Frame) LocalVars() LocalVars {
	return receiver.localVars
}

func (receiver *Frame) OperandStack() *OperandStack {
	return receiver.operandStack
}

func (receiver *Frame) SetNextPC(nextPC int) {
	receiver.nextPC = nextPC
}

func (receiver *Frame) NextPC() int {
	return receiver.nextPC
}

func (receiver *Frame) Thread() *Thread {
	return receiver.thread
}

func (receiver *Frame) Method() *heap.Method {
	return receiver.method
}
