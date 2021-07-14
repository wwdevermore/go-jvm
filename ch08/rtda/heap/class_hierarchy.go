package heap

func (receiver *Class) isAssignableFrom(other *Class) bool {
	s, t := other, receiver
	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.isSubClassOf(t)
			} else {
				return s.isImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJ1Object()
			} else {
				return t.isSubInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJ1Object()
			} else {
				return t.isJ1Cloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
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
