package mavlink1

import (
	"bytes"
	"context"
	"errors"
	"log"
	"time"

	"github.com/mgr9525/go-mavlink1/util"
)

type GetMsgFun func(msg *Mavlink1Msg)

type Mavlink1 struct {
	seq    byte
	sysid  byte
	compid byte

	ctx  context.Context
	cncl context.CancelFunc

	getFun  GetMsgFun
	recvBuf *util.CircleByteBuffer
}
type Mavlink1Msg struct {
	Length byte
	Seq    byte
	Sysid  byte
	Compid byte
	Msgid  byte

	Payload []byte
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
	e.getFun = getfun
	e.ctx, e.cncl = context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-e.ctx.Done():
				e.cncl = nil
				return
			default:
				e.run()
				time.Sleep(time.Millisecond)
			}
		}
	}()
	return nil
}

func (e *Mavlink1) Stopd() bool {
	return e.cncl == nil
}
func (e *Mavlink1) Stop() {
	e.getFun = nil
	if e.cncl != nil {
		e.cncl()
	}
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
	c1, c2 := util.GetChecksumCRC(bthd, btbd)

	sumth := (int(int(crc1)<<8) & 0xffff) | (int(crc2) & 0xff)
	summe := (int(int(c1)<<8) & 0xffff) | (int(c2) & 0xff)
	if sumth != summe {
		log.Printf("checksum error,msgid:%d,th:%d,me:%d!\n", msgid, sumth, summe)
		return
	}

	msg := new(Mavlink1Msg)
	msg.Length = ln
	msg.Seq = seq
	msg.Sysid = sysid
	msg.Compid = compid
	msg.Msgid = msgid
	msg.Payload = btbd
	if e.getFun != nil {
		go e.getFun(msg)
	}
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
		msg.Payload = data
	}
	return msg
}

func (msg *Mavlink1Msg) GetBytes() []byte {
	buf := &bytes.Buffer{}
	bthd := []byte{0xfe, msg.Length, msg.Seq, msg.Sysid, msg.Compid, msg.Msgid}
	c1, c2 := util.GetChecksumCRC(bthd, msg.Payload)
	buf.Write(bthd)
	buf.Write(msg.Payload)
	buf.Write([]byte{c1, c2})
	return buf.Bytes()
}
