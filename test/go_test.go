package main

import (
	"fmt"
	"github.com/mgr9525/go-mavlink1/util"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	s := int16(0x1234)
	b := int8(s)
	fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	return 0x34 == b                            /* {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}*/
}
func Test1(t *testing.T) {
	println("IsLittleEndian:", IsLittleEndian())
	println("val1:", fmt.Sprintf("%f", float32(12.5)))
	val := float32(12.5)
	println("val1:", fmt.Sprintf("%f", val))
	bts := util.LitFloatToByte32(val)
	println(fmt.Sprintf("bts:%x", bts))
	println("val3:", fmt.Sprintf("%f", util.LitByteToFloat32(bts)))
}
