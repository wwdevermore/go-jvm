package heap

import "go-jvm/ch08/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (receiver *MemberRef) copyMemberInfo(refInfo *classfile.ConstantMemberrefInfo) {
	receiver.className = refInfo.ClassName()
	receiver.name, receiver.descriptor = refInfo.NameAndType()
}
