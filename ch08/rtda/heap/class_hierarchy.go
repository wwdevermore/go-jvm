package heap

func (receiver *Class) isAssignableFrom(other *Class) bool {
	s, t := other, receiver
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (receiver *Class) isSubClassOf(other *Class) bool {
	for c := receiver.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (receiver *Class) isImplements(iface *Class) bool {
	for c := receiver; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (receiver *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range receiver.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
