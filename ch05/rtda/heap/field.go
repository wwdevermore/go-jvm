package heap

import "go-jvm/ch05/classfile"

type Field struct {
	ClassMember
	slotId uint
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

func (self *Field) IsStatic() bool {
	return self.accessFlags == ACC_STATIC
}

func (self *Field) IsFinal() bool {
	return self.accessFlags == ACC_FINAL
}

func (self *Field) IsLongOrDouble() bool {
	return self.descriptor == "D" || self.descriptor == "J"
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) isAccessibleTo(other *Class) bool{
	if self.IsPublic(){
		return true
	}
	c := self.class
	if self.IsProtected() {
		return c == other || other.isSubClassOf(c) || c.getPackageName() == other.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == other.getPackageName()
	}
	return other == c
}

func (self *Field) IsPublic() bool {
	return self.accessFlags == ACC_PUBLIC
}

func (self *Field) IsProtected() bool {
	return self.accessFlags == ACC_PROTECTED
}

func (self *Field) IsPrivate() bool {
	return self.accessFlags == ACC_PRIVATE
}
