package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (receiver *Object) Class() *Class {
	return receiver.class
}

func (receiver *Object) Fields() Slots {
	return receiver.fields
}

func (receiver *Object) IsInstanceOf(other *Class) bool {
	return other.isAssignableFrom(receiver.class)
}
