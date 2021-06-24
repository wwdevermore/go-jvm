package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

type LocalVariableTableAttribute struct {
	lineNumberTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc          uint16
	length           uint16
	name_index       uint16
	descriptor_index uint16
	index            uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LocalVariableTableEntry{
			startPc:          reader.readUint16(),
			length:           reader.readUint16(),
			name_index:       reader.readUint16(),
			descriptor_index: reader.readUint16(),
			index:            reader.readUint16(),
		}
	}
}
