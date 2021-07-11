package base

import (
	"go-jvm/ch05/rtda"
	"go-jvm/ch05/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduledClient(thread, class)
	initSuperClass(thread, class)
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {

}

func scheduledClient(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		//exec clinit
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
