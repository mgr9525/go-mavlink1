package main

import (
	"fmt"
	"github.com/mgr9525/go-mavlink1"
	"github.com/mgr9525/go-mavlink1/msg"
)

var mavchan = mavlink1.New()

func main() {
	if mavchan.Start(getMsg) == nil {
		println("mav start!")
		testBytes := []byte{
			0xfe, // stx
			0x10, //Length of payload
			0x35, //Sequence of packet
			0x01, //sysid    ID of message sender system/aircraft
			0x01, //compid  ID of the message sender component
			0x6f, //msgid   ID of message in payload
			0x00, //payload
			0x00, //payload
			0x00, //payload ...
			0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0x7b, 0x5f, 0x1f, 0xba, 0x29, 0x0, 0x00,
			0x35, 0x4f, //crc
		}
		mavchan.Puts(testBytes)
	}

	fmt.Println("input any to exit!")
	fmt.Scanf("%s")
	mavchan.Stop()
}

func getMsg(msg *mavlink1.Mavlink1Msg) {
	fmt.Printf("getMsg from:%x-%x, msgid:%x\n", msg.Sysid, msg.Compid, msg.Msgid)

	sct := new(messages.SetMode)
	sct.TargetSystem = 1
	sct.CustomMode = 1
	sct.BaseMode = 1
	replyMsg := mavchan.NewMsg(1, 2, messages.MSG_ID_SET_MODE, messages.SetMode2Byte(sct))
	replyBytes := mavlink1.GetMsgBytes(replyMsg).Bytes()
	fmt.Print("replyBytes:")
	for _, v := range replyBytes {
		fmt.Printf("%x, ", v)
	}
}
