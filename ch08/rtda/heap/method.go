package heap

import "go-jvm/ch08/classfile"

type Method struct {
	ClassMember
	maxLocals    uint
	maxStack     uint
	code         []byte
	class        *Class
	argSlotCount uint
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "D" || paramType == "J" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++ //this instruction
	}
}

func (receiver *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		receiver.maxLocals = codeAttr.MaxLocals()
		receiver.maxStack = codeAttr.MaxStacks()
		receiver.code = codeAttr.Code()
	}
}

func (receiver *Method) MaxStack() uint {
	return receiver.maxStack
}

func (receiver *Method) MaxLocals() uint {
	return receiver.maxLocals
}

func (receiver *Method) Class() *Class {
	return receiver.class
}

func (receiver *Method) Code() []byte {
	return receiver.code
}

func (receiver *Method) ArgSlotCount() uint {
	return receiver.argSlotCount
}

func (receiver *Method) IsAbstract() bool {
	return receiver.accessFlags&ACC_ABSTRACT == ACC_ABSTRACT
}

func (receiver *Method) IsStatic() bool {
	return receiver.accessFlags&ACC_STATIC == ACC_STATIC
}

func (receiver *Method) IsPublic() bool {
	return receiver.accessFlags&ACC_PUBLIC == ACC_PUBLIC
}

func (receiver *Method) IsProtected() bool {
	return receiver.accessFlags&ACC_PROTECTED == ACC_PROTECTED
}

func (receiver *Method) IsPrivate() bool {
	return receiver.accessFlags&ACC_PRIVATE == ACC_PRIVATE
}

func (receiver *Method) isAccessibleTo(other *Class) bool {
	if receiver.IsPublic() {
		return true
	}
	c := receiver.class
	if receiver.IsProtected() {
		return c == other || other.isSubClassOf(c) || c.GetPackageName() == other.GetPackageName()
	}
	if !receiver.IsPrivate() {
		return c.GetPackageName() == other.GetPackageName()
	}
	return other == c
}
