package heap

import "go-jvm/ch05/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newField(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (receiver *Field) IsStatic() bool {
	return receiver.accessFlags&ACC_STATIC == ACC_STATIC
}

func (receiver *Field) IsFinal() bool {
	return receiver.accessFlags&ACC_FINAL == ACC_FINAL
}

func (receiver *Field) IsLongOrDouble() bool {
	return receiver.descriptor == "D" || receiver.descriptor == "J"
}

func (receiver *Field) SlotId() uint {
	return receiver.slotId
}

func (receiver *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		receiver.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (receiver *Field) ConstValueIndex() uint {
	return receiver.constValueIndex
}

func (receiver *Field) isAccessibleTo(other *Class) bool {
	if receiver.IsPublic() {
		return true
	}
	c := receiver.Class()
	if receiver.IsProtected() {
		return c == other || other.isSubClassOf(c) || c.GetPackageName() == other.GetPackageName()
	}
	if !receiver.IsPrivate() {
		return c.GetPackageName() == other.GetPackageName()
	}
	return other == c
}

func (receiver *Field) IsPublic() bool {
	return receiver.accessFlags&ACC_PUBLIC == ACC_PUBLIC
}

func (receiver *Field) IsProtected() bool {
	return receiver.accessFlags&ACC_PROTECTED == ACC_PROTECTED
}

func (receiver *Field) IsPrivate() bool {
	return receiver.accessFlags&ACC_PRIVATE == ACC_PRIVATE
}
