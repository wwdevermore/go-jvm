package heap

import "go-jvm/ch05/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.cp = cp
	methodRef.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	return methodRef
}

func (receiver *MethodRef) Name() string {
	return receiver.name
}

func (receiver *MemberRef) Descriptor() string {
	return receiver.descriptor
}
