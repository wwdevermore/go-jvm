package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (receiver *Object) Class() *Class {
	return receiver.class
}

func (receiver *Object) Fields() Slots {
	return receiver.data.(Slots)
}

func (receiver *Object) IsInstanceOf(other *Class) bool {
	return other.isAssignableFrom(receiver.class)
}

func (receiver *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field := receiver.class.getField(name, descriptor, false)
	slots := receiver.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (receiver *Object) GetRefVar(name, descriptor string) *Object {
	field := receiver.class.getField(name, descriptor, false)
	slots := receiver.data.(Slots)
	return slots.GetRef(field.slotId)
}
