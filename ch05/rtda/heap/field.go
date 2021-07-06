package heap

import "go-jvm/ch05/classfile"

type Field struct {
	ClassMember
	slotId uint
}

func newField(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (self *Field) IsStatic() bool {
	return self.accessFlags == ACC_STATIC
}

func (self *Field) IsFinal() bool {
	return self.accessFlags == ACC_FINAL
}

func (self *Field) IsLongOrDouble() bool {
	return self.descriptor == "D" || self.descriptor == "J"
}
