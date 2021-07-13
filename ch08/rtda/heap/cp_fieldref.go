package heap

import "go-jvm/ch08/classfile"

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

func (receiver *FieldRef) ResolvedField() *Field {
	if receiver.field == nil {
		receiver.resolveFieldRef()
	}
	return receiver.field
}

func (receiver *FieldRef) resolveFieldRef() {
	d := receiver.cp.class
	c := receiver.ResolveClass()
	field := lookupField(c, receiver.name, receiver.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	receiver.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
