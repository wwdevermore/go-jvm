package heap

import (
	"fmt"
	"go-jvm/ch08/classfile"
	"go-jvm/ch08/classpath"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	classMap    map[string]*Class //loaded classes
	verboseFlag bool
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		classMap:    make(map[string]*Class),
		verboseFlag: verboseFlag,
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	class := self.classMap[name]
	if class != nil {
		return class
	}
	if name[0] == '[' {
		return self.loadArrayClass(name)
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //todo
		name:        name,
		loader:      self,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (receiver *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = receiver
	resolveSuperClass(class)
	resolveInterfaces(class)
	receiver.classMap[class.name] = class
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
	if class.superClassName != "java/lang/Object" && class.superClassName != "" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	l := len(class.InterfaceNames())
	if l > 0 {
		class.interfaces = make([]*Class, l)
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
	calculateStaticFieldSlotIds(class)
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
	class.instanceSlotCount = slotId
}

func calculateStaticFieldSlotIds(class *Class) {
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
	class.staticSlotCount = slotId
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
	switch field.descriptor {
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
		goStr := cp.GetConstant(cpIndex).(string)
		jStr := JString(class.ClassLoader(), goStr)
		vars.SetRef(slotId, jStr)
	}
}
