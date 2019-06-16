package mavlink1

// #include "checksum.h"
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mgr9525/go-mavlink1/util"
	"time"
)

type GetMsgFun func(msg *Mavlink1Msg)

type Mavlink1 struct {
	seq    byte
	sysid  byte
	compid byte

	isRun bool

	getFun  GetMsgFun
	recvBuf *util.CircleByteBuffer
}
type Mavlink1Msg struct {
	Length byte
	Seq    byte
	Sysid  byte
	Compid byte
	Msgid  byte

	Payload bytes.Buffer
}

func New() *Mavlink1 {
	e := new(Mavlink1)
	e.seq = 0
	e.recvBuf = util.NewCircleByteBuffer(1024 * 5)
	return e
}

func (e *Mavlink1) Start(getfun GetMsgFun) error {
	if getfun == nil {
		return errors.New("getfun is nil")
	}
	e.isRun = true
	e.getFun = getfun
	go func() {
		for e.isRun && e.getFun != nil {
			e.run()
			time.Sleep(time.Millisecond)
		}
	}()
	return nil
}

func (e *Mavlink1) Stop() {
	e.isRun = false
	e.getFun = nil
}

func (e *Mavlink1) run() {
	defer util.Recovers("run")

	lens := e.recvBuf.GetLen()
	for i := 0; i < lens; i++ {
		if e.recvBuf.Geti(0) != 0xfe {
			e.recvBuf.GetByte()
		} else {
			break
		}
	}

	if e.recvBuf.GetLen() < 6 {
		return
	}

	ln := e.recvBuf.Geti(1)
	seq := e.recvBuf.Geti(2)
	sysid := e.recvBuf.Geti(3)
	compid := e.recvBuf.Geti(4)
	msgid := e.recvBuf.Geti(5)

	length := int(ln & 0xff)
	if e.recvBuf.GetLen() < length+8 {
		return
	}

	//------head
	bthd := make([]byte, 6)
	e.recvBuf.Read(bthd)
	//------head end

	//------body
	btbd := make([]byte, length)
	e.recvBuf.Read(btbd)
	//------body end

	crc1, _ := e.recvBuf.GetByte()
	crc2, _ := e.recvBuf.GetByte()
	c1, c2 := GetChecksumCRC(&bthd, &btbd)

	sumth := (int(int(crc1)<<8) & 0xffff) | (int(crc2) & 0xff)
	summe := (int(int(c1)<<8) & 0xffff) | (int(c2) & 0xff)
	if sumth != summe {
		fmt.Printf("checksum error,msgid:%d,th:%d,me:%d!\n", msgid, sumth, summe)
		return
	}

	msg := new(Mavlink1Msg)
	msg.Length = ln
	msg.Seq = seq
	msg.Sysid = sysid
	msg.Compid = compid
	msg.Msgid = msgid
	msg.Payload.Write(btbd)
	go e.getFun(msg)
}

func (e *Mavlink1) Puts(bts []byte) (int, error) {
	return e.recvBuf.Write(bts)
}
func (e *Mavlink1) NextSeq() byte {
	ret := e.seq
	if e.seq >= 255 {
		e.seq = 0
	} else {
		e.seq += 1
	}
	return ret
}
func (e *Mavlink1) NewMsg(sysid, compid, msgid int, data []byte) *Mavlink1Msg {
	msg := new(Mavlink1Msg)
	msg.Seq = e.NextSeq()
	msg.Sysid = byte(sysid)
	msg.Compid = byte(compid)
	msg.Msgid = byte(msgid)
	msg.Length = byte(len(data))
	if data != nil {
		msg.Payload.Write(data)
	}
	return msg
}

func GetMsgBytes(msg *Mavlink1Msg) *bytes.Buffer {
	buf := new(bytes.Buffer)
	bthd := []byte{0xfe, msg.Length, msg.Seq, msg.Sysid, msg.Compid, msg.Msgid}
	btbd := msg.Payload.Bytes()
	c1, c2 := GetChecksumCRC(&bthd, &btbd)
	buf.Write(bthd)
	buf.Write(btbd)
	buf.Write([]byte{c1, c2})
	return buf
}

func GetChecksumCRC(bthd, btbd *[]byte) (byte, byte) {
	bthds := *bthd
	btbds := *btbd
	checksum := C.crc_calculate(C.CBytes(bthds[1:]), 5)
	C.crc_accumulate_buffer(&checksum, C.CBytes(btbds), C.int(len(btbds)))
	C.crc_accumulate(C.uchar(MSG_CRC[bthds[5]]&0xff), &checksum)
	c1 := byte(checksum & 0xff)
	c2 := byte((checksum >> 8) & 0xff)
	return c1, c2
}
