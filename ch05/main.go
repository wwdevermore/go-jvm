package main

import (
	"fmt"
	"go-jvm/ch05/classpath"
	"go-jvm/ch05/interpreter"
	"go-jvm/ch05/rtda/heap"
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
	classLoader := heap.NewClassLoader(cp)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetStaticMethod("main", "([Ljava/lang/String;)V")
	if mainMethod != nil {
		interpreter.Interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
