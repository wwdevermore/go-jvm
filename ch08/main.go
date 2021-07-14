package main

import (
	"fmt"
	"go-jvm/ch08/classpath"
	"go-jvm/ch08/interpreter"
	"go-jvm/ch08/rtda/heap"
)

func main() {
	//cmd := parseCmd()
	//fmt.Println("version 0.0.1")
	startJVM()
}

func startJVM() {
	cp := classpath.Parse("C:\\Program Files\\Java\\jdk1.8.0_121\\jre\\", "")
	className := "BbSortTEST"
	classLoader := heap.NewClassLoader(cp, true)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetStaticMethod("main", "([Ljava/lang/String;)V")
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", "InvokeDemo.class")
	}
}
