package main

import (
	"fmt"
	"github.com/mgr9525/go-mavlink1"
)

func main() {
	var mavchan = mavlink1.New()
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

	replyMsg := new(mavlink1.Mavlink1Msg)
	replyMsg.Length = 2
	replyMsg.Seq = 1
	replyMsg.Sysid = 2
	replyMsg.Compid = 1
	replyMsg.Msgid = mavlink1.SET_MODE //test
	replyMsg.Payload.Write([]byte{0x11, 0x22})
	replyBytes := mavlink1.GetMsgBytes(replyMsg).Bytes()
	fmt.Print("replyBytes:")
	for _, v := range replyBytes {
		fmt.Printf("%x, ", v)
	}
}
