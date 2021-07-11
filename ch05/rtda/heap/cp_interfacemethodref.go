package heap

import "go-jvm/ch05/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (receiver *InterfaceMethodRef) Name() string{
	return receiver.name
}

func (receiver *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if receiver.method == nil {
		receiver.resolveInterfaceMethodRef()
	}
	return receiver.method
}

func (receiver *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := receiver.cp.class
	c := receiver.ResolveClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, receiver.name, receiver.descriptor)
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessibleError")
	}
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	receiver.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
