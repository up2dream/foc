package incr

type Pid struct {
	TargetValue float64
	ActualValue float64
	Err         float64
	ErrNext     float64
	ErrLast     float64
	Kp, Ki, Kd  float64
}

func (pid *Pid) Init() {
	pid.TargetValue = 0.0
	pid.ActualValue = 0.0
	pid.Err = 0.0
	pid.ErrLast = 0.0
	pid.ErrNext = 0.0
	pid.Kp = 0.21
	pid.Ki = 0.80
	pid.Kd = 0.01
}

func (pid *Pid) Realize(tempValue float64) float64 {
	pid.Err = pid.TargetValue - tempValue
	incrementVal := pid.Kp*(pid.Err-pid.ErrNext) + pid.Ki*pid.Err + pid.Kd*(pid.Err-2*pid.ErrNext+pid.ErrLast)
	pid.ActualValue += incrementVal
	pid.ErrLast = pid.ErrNext
	pid.ErrNext = pid.Err
	return pid.ActualValue
}
