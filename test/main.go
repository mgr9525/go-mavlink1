package main

import (
	"encoding/hex"
	"fmt"
	"github.com/mgr9525/go-mavlink1"
	"github.com/mgr9525/go-mavlink1/messages"
	"github.com/mgr9525/go-mavlink1/util"
	"time"
	"unsafe"
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

		time.Sleep(time.Second)

		sct := &messages.Attitude{}
		copy(sct.TimeBootMs[:], util.LitIntToByte(5, len(sct.TimeBootMs)))
		copy(sct.Roll[:], util.LitFloatToByte32(12.5))
		copy(sct.Pitch[:], util.LitFloatToByte32(16.4))
		copy(sct.Yaw[:], util.LitFloatToByte32(19.3))
		copy(sct.Rollspeed[:], util.LitFloatToByte32(321.3))
		copy(sct.Pitchspeed[:], util.LitFloatToByte32(57.8))
		copy(sct.Yawspeed[:], util.LitFloatToByte32(21.7))
		payload, err := util.Struct2Bytes(unsafe.Pointer(sct), unsafe.Sizeof(*sct))
		if err == nil {
			replyMsg := mavchan.NewMsg(1, 2, messages.MSG_ID_ATTITUDE, payload)
			mavchan.Puts(replyMsg.GetBytes())
		}
	}

	fmt.Println("input any to exit!")
	time.Sleep(time.Second * 2)
	mavchan.Stop()
}

func getMsg(msg *mavlink1.Mavlink1Msg) {
	fmt.Printf("getMsg from:%x-%x, msgid:%x\n", msg.Sysid, msg.Compid, msg.Msgid)
	if msg.Msgid == messages.MSG_ID_ATTITUDE {
		atu := &messages.Attitude{}
		err := util.Bytes2Struct(unsafe.Pointer(atu), msg.Payload, unsafe.Sizeof(*atu))
		if err != nil {
			println("convert payload to msg err:", err)
			return
		}
		println("get ATTITUDE:", hex.EncodeToString(msg.Payload))
		println(fmt.Sprintf("roll:%f,pitch:%f,yaw:%f",
			util.LitByteToFloat32(atu.Roll[:]), util.LitByteToFloat32(atu.Pitch[:]), util.LitByteToFloat32(atu.Yaw[:])))
		println(fmt.Sprintf("rollspeed:%f,pitchspeed:%f,yawspeed:%f",
			util.LitByteToFloat32(atu.Rollspeed[:]), util.LitByteToFloat32(atu.Pitchspeed[:]), util.LitByteToFloat32(atu.Yawspeed[:])))
	} else {
		sct := &messages.SetMode{}
		sct.TargetSystem = 1
		sct.BaseMode = 1
		copy(sct.CustomMode[:], util.LitIntToByte(messages.MAV_MODE_FLAG_DECODE_POSITION_GUIDED, len(sct.CustomMode)))
		payload, err := util.Struct2Bytes(unsafe.Pointer(sct), unsafe.Sizeof(*sct))
		if err != nil {
			println("convert msg to payload err:", err)
			return
		}
		replyMsg := mavchan.NewMsg(1, 2, messages.MSG_ID_SET_MODE, payload)
		println("replyBytes:", hex.EncodeToString(replyMsg.GetBytes()))
	}
}
