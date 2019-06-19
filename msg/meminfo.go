package messages

import "unsafe"

const MSG_ID_MEMINFO = 152

type Meminfo struct {
	Brkval    uint16 /*<  heap top*/
	Freemem   uint16 /*< [bytes] free memory*/
	Freemem32 uint32 /*< [bytes] free memory (32 bit)*/
}

func Byte2Meminfo(bts *[]byte) *Meminfo {
	sct := *(**Meminfo)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func Meminfo2Byte(sct *Meminfo) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
