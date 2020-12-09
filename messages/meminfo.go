package messages

const MSG_ID_MEMINFO = 152

type Meminfo struct {
	Brkval    [2]byte /* uint16 /*<  heap top*/
	Freemem   [2]byte /* uint16 /*< [bytes] free memory*/
	Freemem32 [4]byte /* uint32 /*< [bytes] free memory (32 bit)*/
}
