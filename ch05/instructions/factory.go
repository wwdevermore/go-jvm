package instructions

import (
	"fmt"
	"go-jvm/ch05/instructions/base"
	. "go-jvm/ch05/instructions/comparisons"
	. "go-jvm/ch05/instructions/constants"
	. "go-jvm/ch05/instructions/control"
	. "go-jvm/ch05/instructions/conversions"
	. "go-jvm/ch05/instructions/extended"
	. "go-jvm/ch05/instructions/loads"
	. "go-jvm/ch05/instructions/math"
	. "go-jvm/ch05/instructions/references"
	. "go-jvm/ch05/instructions/stack"
	. "go-jvm/ch05/instructions/stores"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return bi_push
	case 0x11:
		return si_push
	case 0x12:
		return ldc
	case 0x13:
		return ldc_w
	case 0x14:
		return ldc2_w
	case 0x15:
		return iload
	case 0x16:
		return lload
	case 0x17:
		return fload
	case 0x18:
		return dload
	case 0x19:
		return aload
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x36:
		return istore
	case 0x37:
		return lstore
	case 0x38:
		return fstore
	case 0x39:
		return dstore
	case 0x3a:
		return astore
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3

	case 0x57:
		return pop
	case 0x58:
		return pop_2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap

	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x65:
		return lsub

	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem

	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land

	case 0x84:
		return iinc
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f

	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return ifeq
	case 0x9a:
		return ifne
	case 0x9b:
		return iflt
	case 0x9c:
		return ifge
	case 0x9d:
		return ifgt
	case 0x9e:
		return ifle
	case 0x9f:
		return if_icmpeq
	case 0xa0:
		return if_icmpne
	case 0xa1:
		return if_icmplt
	case 0xa2:
		return if_icmpge
	case 0xa3:
		return if_icmpgt
	case 0xa4:
		return if_icmple
	case 0xa5:
		return if_acmpeq
	case 0xa6:
		return if_acmpne
	case 0xa7:
		return go_to

	case 0xaa:
		return tableswitch
	case 0xab:
		return lookupswitch
	case 0xac:
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn

	case 0xb1:
		return _return
	case 0xb2:
		return getstatic
	case 0xb3:
		return putstatic
	case 0xb4:
		return getfield
	case 0xb5:
		return putfield
	case 0xb6:
		return invoke_virtual
	case 0xb7:
		return invoke_special
	case 0xb8:
		return invoke_static
	case 0xb9:
		return invoke_interface
	case 0xbb:
		return _new

	case 0xc0:
		return checkcast
	case 0xc1:
		return instanceof
	case 0xc4:
		return wide
	case 0xc6:
		return ifnull
	case 0xc7:
		return ifnonnull
	case 0xc8:
		return goto_w

	default:
		panic(fmt.Errorf("java unsupport opcode: HEX:0x%x, DEC: %d", opcode, opcode))
	}
}

var (
	nop         = &NOP{}
	aconst_null = &ACONST_NULL{}
	iconst_m1   = &ICONST_M1{}
	iconst_0    = &ICONST_0{}
	iconst_1    = &ICONST_1{}
	iconst_2    = &ICONST_2{}
	iconst_3    = &ICONST_3{}
	iconst_4    = &ICONST_4{}
	iconst_5    = &ICONST_5{}
	lconst_0    = &LCONST_0{}
	lconst_1    = &LCONST_1{}
	fconst_0    = &FCONST_0{}
	fconst_1    = &FCONST_1{}
	fconst_2    = &FCONST_2{}
	dconst_0    = &DCONST_0{}
	dconst_1    = &DCONST_1{}
	bi_push     = &BIPUSH{}
	si_push     = &SIPUSH{}
	//idc = &
	iload   = &ILOAD{}
	lload   = &LLOAD{}
	fload   = &FLOAD{}
	dload   = &DLOAD{}
	aload   = &ALOAD{}
	iload_0 = &ILOAD_0{}
	iload_1 = &ILOAD_1{}
	iload_2 = &ILOAD_2{}
	iload_3 = &ILOAD_3{}
	lload_0 = &LLOAD_0{}
	lload_1 = &LLOAD_1{}
	lload_2 = &LLOAD_2{}
	lload_3 = &LLOAD_3{}
	fload_0 = &FLOAD_0{}
	fload_1 = &FLOAD_1{}
	fload_2 = &FLOAD_2{}
	fload_3 = &FLOAD_3{}
	dload_0 = &DLOAD_0{}
	dload_1 = &DLOAD_1{}
	dload_2 = &DLOAD_2{}
	dload_3 = &DLOAD_3{}
	aload_0 = &ALOAD_0{}
	aload_1 = &ALOAD_1{}
	aload_2 = &ALOAD_2{}
	aload_3 = &ALOAD_3{}
	//iaload = &IALOAD{}
	istore   = &ISTORE{}
	lstore   = &LSTORE{}
	fstore   = &FSTORE{}
	dstore   = &DSTORE{}
	astore   = &ASTORE{}
	istore_0 = &ISTORE_0{}
	istore_1 = &ISTORE_1{}
	istore_2 = &ISTORE_2{}
	istore_3 = &ISTORE_3{}
	lstore_0 = &LSTORE_0{}
	lstore_1 = &LSTORE_1{}
	lstore_2 = &LSTORE_2{}
	lstore_3 = &LSTORE_3{}
	fstore_0 = &FSTORE_0{}
	fstore_1 = &FSTORE_1{}
	fstore_2 = &FSTORE_2{}
	fstore_3 = &FSTORE_3{}
	dstore_0 = &DSTORE_0{}
	dstore_1 = &DSTORE_1{}
	dstore_2 = &DSTORE_2{}
	dstore_3 = &DSTORE_3{}
	astore_0 = &ASTORE_0{}
	astore_1 = &ASTORE_1{}
	astore_2 = &ASTORE_2{}
	astore_3 = &ASTORE_3{}
	//iastore = &IASTORE{}
	pop     = &POP{}
	pop_2   = &POP2{}
	dup     = &DUP{}
	dup_x1  = &DUP_X1{}
	dup_x2  = &DUP_X2{}
	dup2    = &DUP2{}
	dup2_x1 = &DUP2_X1{}
	dup2_x2 = &DUP2_X2{}
	swap    = &SWAP{}

	irem = &IREM{}
	lrem = &LREM{}
	frem = &FREM{}
	drem = &DREM{}

	ishl  = &ISHL{}
	lshl  = &LSHL{}
	ishr  = &ISHR{}
	lshr  = &LSHR{}
	iushr = &IUSHR{}
	lushr = &LUSHR{}
	iand  = &IAND{}
	land  = &LAND{}

	iinc = &IINC{}
	i2l  = &I2L{}
	i2f  = &I2F{}
	i2d  = &I2D{}
	l2i  = &L2I{}
	l2f  = &L2F{}
	l2d  = &L2D{}
	f2i  = &F2I{}
	f2l  = &F2L{}
	f2d  = &F2D{}
	d2i  = &D2I{}
	d2l  = &D2L{}
	d2f  = &D2F{}

	lcmp      = &LCMP{}
	fcmpl     = &FCMPL{}
	fcmpg     = &FCMPG{}
	dcmpl     = &DCMPL{}
	dcmpg     = &DCMPG{}
	ifeq      = &IFEQ{}
	ifne      = &IFNE{}
	iflt      = &IFLT{}
	ifge      = &IFGE{}
	ifgt      = &IFGT{}
	ifle      = &IFLE{}
	if_icmpeq = &IF_ICMPEQ{}
	if_icmpne = &IF_ICMPNE{}
	if_icmplt = &IF_ICMPLT{}
	if_icmpge = &IF_ICMPGE{}
	if_icmpgt = &IF_ICMPGT{}
	if_icmple = &IF_ICMPLE{}
	if_acmpeq = &IF_ACMPEQ{}
	if_acmpne = &IF_ACMPNE{}
	go_to     = &GOTO{}

	iadd = &IADD{}
	ladd = &LADD{}

	lsub = &LSUB{}

	tableswitch  = &TABLE_SWITCH{}
	lookupswitch = &LOOKUP_SWITCH{}
	ldc          = &LDC{}
	ldc_w        = &LDC_W{}
	ldc2_w       = &LDC2_W{}
	getstatic    = &GET_STATIC{}
	putstatic    = &PUT_STATIC{}
	getfield     = &GET_FIELD{}
	putfield     = &PUT_FIELD{}
	checkcast    = &CHECK_CAST{}
	instanceof   = &INSTANCE_OF{}

	ifnull    = &IFNULL{}
	ifnonnull = &IFNONNULL{}
	goto_w    = &GOTO_W{}
	wide      = &WIDE{}
	_new       = &NEW{}

	invoke_special = &INVOKE_SPECIAL{}
	invoke_virtual = &INVOKE_VIRTUAL{}
	invoke_static = &INVOKE_STATIC{}
	invoke_interface = &INVOKE_INTERFACE{}

	_return = &RETURN{}
	ireturn = &IRETURN{}
	lreturn = &LRETURN{}
	dreturn = &DRETURN{}
	freturn = &FRETURN{}
	areturn = &ARETURN{}
)
