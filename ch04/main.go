package main

import (
	"fmt"
	"go-jvm/ch04/rtda"
)

func main() {
	cmd := parseCmd()
	fmt.Println("version 0.0.1")
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(localVars rtda.LocalVars) {
	println("======testLocalVars======")
	localVars.SetInt(0, 111)
	localVars.SetInt(1, 123)
	localVars.SetLong(2, 2997924580)
	localVars.SetLong(4, -2997924580)
	localVars.SetFloat(6, 3.1415926)
	localVars.SetDouble(7, 2.71826485845)
	localVars.SetRef(9, nil)
	println(localVars.GetInt(0))
	println(localVars.GetInt(1))
	println(localVars.GetLong(2))
	println(localVars.GetFloat(6))
	println(localVars.GetDouble(7))
	println(localVars.GetRef(9))
}

func testOperandStack(operandStack *rtda.OperandStack) {
	println("======testOperandStack======")
	operandStack.PushDouble(1.33)
	println(operandStack.PopDouble())
}
