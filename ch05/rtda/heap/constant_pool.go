package heap

import (
	"fmt"
	"go-jvm/ch05/classfile"
)

type ConstantPool struct {
	class *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool{
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	runtimeConstantsPool := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {

		}
	}
	return runtimeConstantsPool
}

func (self *ConstantPool) GetConstant(index uint) Constant{
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
