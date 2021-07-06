package heap

import "go-jvm/ch05/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (receiver *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	receiver.accessFlags = memberInfo.Accessflags()
	receiver.name = memberInfo.Name()
	receiver.descriptor = memberInfo.Descriptor()
}

func (receiver *ClassMember) Class() *Class {
	return receiver.class
}

func (receiver *ClassMember) Name() string {
	return receiver.name
}

func (receiver *ClassMember) Descriptor() string {
	return receiver.descriptor
}