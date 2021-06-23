package classfile

type ConstantMemberInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberInfo) class() string{
	return self.cp.getUtf8(self.classIndex)
}

func (self *ConstantMemberInfo) nameAndType() string{
	return self.cp.getUtf8(self.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct { ConstantMemberInfo }
type ConstantMethodRefInfo struct { ConstantMemberInfo }
type ConstantInterfaceMethodRefInfo struct { ConstantMemberInfo }