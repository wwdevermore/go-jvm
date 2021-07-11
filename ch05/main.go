package main

import (
	"fmt"
	"go-jvm/ch05/classpath"
	"go-jvm/ch05/interpreter"
	"go-jvm/ch05/rtda/heap"
)

func main() {
	//cmd := parseCmd()
	//fmt.Println("version 0.0.1")
	startJVM()
}

func startJVM() {
	cp := classpath.Parse("/Library/Java/JavaVirtualMachines/jdk1.8.0_291.jdk/Contents/Home/jre", "")
	className := "InvokeDemo"
	classLoader := heap.NewClassLoader(cp, true)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetStaticMethod("main", "([Ljava/lang/String;)V")
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", "InvokeDemo.class")
	}
}
