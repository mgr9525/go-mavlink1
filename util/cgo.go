package util

/*
#include<string.h>
#include "checksum.h"
void mavmemcpy(void *dst,void*src,unsigned long ln){
	memcpy(dst,src,(size_t)ln);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func GetChecksumCRC(bthd, btbd []byte) (byte, byte) {
	bthds := bthd[1:]
	checksum := C.crc_calculate(unsafe.Pointer(&bthds[0]), 5)
	C.crc_accumulate_buffer(&checksum, unsafe.Pointer(&btbd[0]), C.int(len(btbd)))
	C.crc_accumulate(C.uchar(MSG_CRC[bthd[5]]&0xff), &checksum)
	c1 := byte(checksum & 0xff)
	c2 := byte((checksum >> 8) & 0xff)
	return c1, c2
}

type SliceMock struct {
	addr unsafe.Pointer
	len  uint
	cap  uint
}

func Struct2Bytes(pt unsafe.Pointer, ln uintptr) ([]byte, error) {
	mock := &SliceMock{
		addr: pt,
		cap:  uint(ln),
		len:  uint(ln),
	}
	bts := *(*[]byte)(unsafe.Pointer(mock))
	rtbts := make([]byte, mock.len)
	copy(rtbts, bts)
	return rtbts, nil
}
func Bytes2Struct(pt unsafe.Pointer, bts []byte, ln uintptr) error {
	if pt == nil {
		return errors.New("param is nil")
	}
	dataptr := uintptr(pt)
	C.mavmemcpy(unsafe.Pointer(dataptr), unsafe.Pointer(&bts[0]), C.ulong(ln))
	return nil
}
