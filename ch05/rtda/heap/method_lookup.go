package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMethodInInterfaces(iface []*Class, name, descriptor string) *Method {
	for _, iface := range iface {
		for _, m := range iface.methods {
			if m.name == name && m.descriptor == descriptor {
				return m
			}
		}
		m := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if m != nil {
			return m
		}
	}
	return nil
}
