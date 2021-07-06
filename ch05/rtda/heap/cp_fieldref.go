package heap

import "go-jvm/ch05/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	fieldRef := &FieldRef{}
	fieldRef.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	fieldRef.cp = cp
	return fieldRef
}
