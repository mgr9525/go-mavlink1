package messages

const MSG_ID_SENSOR_OFFSETS = 150

type SensorOffsets struct {
	MagDeclination [4]byte /* float32 /*< [rad] magnetic declination (radians)*/
	RawPress       [4]byte /* int32   /*<  raw pressure from barometer*/
	RawTemp        [4]byte /* int32   /*<  raw temperature from barometer*/
	GyroCalX       [4]byte /* float32 /*<  gyro X calibration*/
	GyroCalY       [4]byte /* float32 /*<  gyro Y calibration*/
	GyroCalZ       [4]byte /* float32 /*<  gyro Z calibration*/
	AccelCalX      [4]byte /* float32 /*<  accel X calibration*/
	AccelCalY      [4]byte /* float32 /*<  accel Y calibration*/
	AccelCalZ      [4]byte /* float32 /*<  accel Z calibration*/
	MagOfsX        [2]byte /* int16   /*<  magnetometer X offset*/
	MagOfsY        [2]byte /* int16   /*<  magnetometer Y offset*/
	MagOfsZ        [2]byte /* int16   /*<  magnetometer Z offset*/
}
