package rtda

import "go-jvm/ch05/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
