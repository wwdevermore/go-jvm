package main

import (
	"fmt"
	"go-jvm/ch05/classfile"
	"go-jvm/ch05/classpath"
	"go-jvm/ch05/interpreter"
	"strings"
)

func main() {
	cmd := parseCmd()
	fmt.Println("version 0.0.1")
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != err {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

/*func testLocalVars(localVars rtda.LocalVars) {
	println("======testLocalVars======")
	localVars.SetInt(0, 111)
	localVars.SetInt(1, -123)
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
*/
