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

func (receiver *MethodRef) ResolvedMethod() *Method {
	if receiver.method == nil {
		receiver.resolveMethodRef()
	}
	return receiver.method
}

func (receiver *MethodRef) resolveMethodRef() {
	d := receiver.cp.class
	c := receiver.ResolveClass()
	if c.IsInterface() {
		panic("java.lang.IncompatiableClassChangeError")
	}
	method := lookupMethod(c, receiver.name, receiver.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	receiver.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
