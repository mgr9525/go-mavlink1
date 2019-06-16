package util

import "fmt"

func Recovers(name string) {
	if err := recover(); err != nil {
		fmt.Print("ruisRecover(" + name + "):")
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	}
}

func Bytes2Short(bts []byte) int {
	var ret = 0
	if len(bts) >= 2 {
		ret = (int(int(bts[0])<<8) & 0xffff) | (int(bts[1]) & 0xff)
	}
	return ret
}
func Short2Bytes(v int) []byte {
	var bts = []byte{0, 0}
	bts[0] = byte((v >> 8) & 0xff)
	bts[1] = byte(v & 0xff)
	return bts
}
