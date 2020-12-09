# Golang Mavlink1.0

This is a very stable mavlink message parsing library.It doesn't care how your data is transmitted. more se [mavlink.io](https://mavlink.io/en/#supported_languages)

## How to use
```
go get github.com/mgr9525/go-mavlink1
```

## Begin
```
mavchan := mavlink1.New()
mavchan.Start(getMsg)
mavchan.Puts(mavData)
```
> getMsg（func）: It's a function call when get message

> mavData([]byte): It's data stream of mavlink.v1,it could receive from serialport or tcp socket.

## More detail see folder test
### [test/main.go](test/main.go "more examples")


## Output
```
mav start!
getMsg from:1-1, msgid:6f
replyBytes: fe060001020b080000000101f1da
input any to exit!
getMsg from:1-2, msgid:1e
get ATTITUDE: 05000000000048413333834166669a4166a6a043333367429a99ad41
roll:12.500000,pitch:16.400000,yaw:19.299999
rollspeed:321.299988,pitchspeed:57.799999,yawspeed:21.700001

```


## How to add message
### More detail see folder "messages"
> [Set fly system mode message!!!](messages/setmode.go "more examples")

> Notice the sizeof of the structure,it control with byte and []byte,con't use int32,float32...,see [Struct sizeof](https://stackoverflow.com/questions/34219232/struct-has-different-size-if-the-field-order-is-different)
