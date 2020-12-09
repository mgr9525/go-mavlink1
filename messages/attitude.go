package messages

const MSG_ID_ATTITUDE = 30

type Attitude struct {
	TimeBootMs [4]byte /*uint32  < [ms] Timestamp (time since system boot).*/
	Roll       [4]byte /*float32 < [rad] Roll angle (-pi..+pi)*/
	Pitch      [4]byte /*float32 < [rad] Pitch angle (-pi..+pi)*/
	Yaw        [4]byte /*float32 < [rad] Yaw angle (-pi..+pi)*/
	Rollspeed  [4]byte /*float32 < [rad/s] Roll angular speed*/
	Pitchspeed [4]byte /*float32 < [rad/s] Pitch angular speed*/
	Yawspeed   [4]byte /*float32 < [rad/s] Yaw angular speed*/
}
