package heap

import "go-jvm/ch05/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.Accessflags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
