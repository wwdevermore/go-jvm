package references

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

//anewarray指令也需要两个操作数。第一个操作数是uint16索
//引，来自字节码。通过这个索引可以从当前类的运行时常量池中找
//到一个类符号引用，解析这个符号引用就可以得到数组元素的类。
//第二个操作数是数组长度，从操作数栈中弹出。Execute（）方法根据
//数组元素的类型和数组长度创建引用类型数
func (receiver *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(receiver.Index).(*heap.ClassRef)
	componentClass := classRef.ResolveClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
