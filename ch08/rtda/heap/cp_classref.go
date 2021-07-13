package heap

import "go-jvm/ch08/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.className = classInfo.Name()
	return classRef
}
