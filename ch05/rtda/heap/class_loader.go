package heap

import (
	"fmt"
	"go-jvm/ch05/classfile"
	"go-jvm/ch05/classpath"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class //loaded classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	class := self.classMap[name]
	if class != nil {
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	classfile, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(classfile)
}

func resolveSuperClass(class *Class) {
	if class.superClassName != "java.lang.Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	len := len(class.InterfaceNames())
	if len > 0 {
		class.interfaces = make([]*Class, len)
		for i, interfaceName := range class.InterfaceNames() {
			if class.loader.classMap[interfaceName] == nil {
				class.interfaces[i] = class.loader.LoadClass(interfaceName)
			}
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//nothing todo
}

func prepare(class *Class) {
	calculateStaticFieldSlotids(class)
	calculateInstanceFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calculateInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
}

func calculateStaticFieldSlotids(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	switch field.descriptor{
	case "C", "B", "Z", "S", "I":
		val := cp.GetConstant(cpIndex).(int32)
		vars.SetInt(slotId, val)
	case "J":
		val := cp.GetConstant(cpIndex).(int64)
		vars.SetLong(slotId, val)
	case "F":
		val := cp.GetConstant(cpIndex).(float32)
		vars.SetFloat(slotId, val)
	case "D":
		val := cp.GetConstant(cpIndex).(float64)
		vars.SetDouble(slotId, val)
	case "Ljava.lang.String":
		panic("unsupport static final field descriptor type")
	}
}
