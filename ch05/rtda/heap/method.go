package heap

import "go-jvm/ch05/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
	class *Class
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxLocals = codeAttr.MaxLocals()
		self.maxStack = codeAttr.MaxStacks()
		self.code = codeAttr.Code()
	}
}

func (self *Method) Class() *Class {
	return self.class
}