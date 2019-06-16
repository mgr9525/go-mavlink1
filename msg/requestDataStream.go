package messages

import "unsafe"

const MSG_ID_REQUEST_DATA_STREAM = 66

type RequestDataStream struct {
	ReqMessageRate  uint16 /*< [Hz] The requested message rate*/
	TargetSystem    uint8  /*<  The target requested to send the message stream.*/
	TargetComponent uint8  /*<  The target requested to send the message stream.*/
	ReqStreamId     uint8  /*<  The ID of the requested data stream*/
	StartStop       uint8  /*<  1 to start sending, 0 to stop sending.*/
}

func Byte2RequestDataStream(bts *[]byte) *RequestDataStream {
	sct := *(**RequestDataStream)(unsafe.Pointer(bts))
	return sct
}
func RequestDataStream2Byte(sct *RequestDataStream) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
