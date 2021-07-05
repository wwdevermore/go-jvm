package heap

import "go-jvm/ch02/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range methods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
}
