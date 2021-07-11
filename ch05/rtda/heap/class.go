package heap

import (
	"go-jvm/ch05/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newField(class, cf.Fields())
	class.methods = newMethod(class, cf.Methods())
	return class
}

func (receiver *Class) StartInit() {
	receiver.initStarted = true
}

func (receiver *Class) InitStarted() bool {
	return receiver.initStarted
}

func (receiver *Class) IsPublic() bool {
	return 0 != receiver.accessFlags&ACC_PUBLIC
}

func (receiver *Class) IsFinal() bool {
	return 0 != receiver.accessFlags&ACC_FINAL
}

func (receiver *Class) IsSynthetic() bool {
	return 0 != receiver.accessFlags&ACC_SYNTHETIC
}

func (receiver *Class) IsPrivate() bool {
	return 0 != receiver.accessFlags&ACC_PRIVATE
}

func (receiver *Class) InterfaceNames() []string {
	return receiver.interfaceNames
}

func (receiver *Class) Fields() []*Field {
	return receiver.fields
}

func (receiver *Class) isAccessibleTo(other *Class) bool {
	return receiver.IsPublic() || receiver.GetPackageName() == other.GetPackageName()
}

func (receiver *Class) GetPackageName() string {
	if i := strings.LastIndex(receiver.name, "/"); i >= 0 {
		return receiver.name[:i]
	}
	return ""
}

func (receiver *Class) ConstantPool() *ConstantPool {
	return receiver.constantPool
}

func (receiver *Class) IsSupper() bool {
	return receiver.accessFlags&ACC_SUPER != 0
}

func (receiver *Class) IsInterface() bool {
	return receiver.accessFlags&ACC_INTERFACE != 0
}

func (receiver *Class) IsAbstract() bool {
	return receiver.accessFlags&ACC_ABSTRACT != 0
}

func (receiver *Class) NewObject() *Object {
	return newObject(receiver)
}

func (receiver *Class) StaticVars() Slots {
	return receiver.staticVars
}

func (receiver *Class) IsSuperClassOf(other *Class) bool {
	return receiver.superClass == other
}

func (receiver *Class) IsSubClassOf(other *Class) bool {
	return other.superClass == receiver
}

func (receiver *Class) IsImplementsOf(other *Class) bool {
	for _, iface := range receiver.interfaces {
		if iface == other {
			return true
		}
	}
	return false
}

func (receiver *Class) Name() string {
	return receiver.name
}

func (receiver *Class) GetStaticMethod(name, descriptor string) *Method {
	return receiver.getStaticMethod(name, descriptor)
}

func (receiver *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range receiver.methods {
		if method.name == name && method.descriptor == descriptor && method.IsStatic() {
			return method
		}
	}
	return nil
}

func (receiver *Class) GetClinitMethod() *Method {
	return receiver.getStaticMethod("<clinit>", "()V")
}

func (receiver *Class) SuperClass() *Class {
	return receiver.superClass
}
