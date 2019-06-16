package messages

import (
	"bytes"
	"encoding/binary"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func Byte2PayloadStruct(t interface{}, bts []byte) error {
	buf := bytes.NewBuffer(bts)
	err := binary.Read(buf, binary.LittleEndian, t)
	if err != nil {
		return err
	}
	return nil
}
func Struct2PayloadByte(t interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, *(t.(*interface{})))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
