package pos

type Pid struct {
	TargetValue float64
	ActualValue float64
	Err         float64
	ErrLast     float64
	Kp, Ki, Kd  float64
	Integral    float64
}

func (pid *Pid) Init() {
	pid.TargetValue = 0.0
	pid.ActualValue = 0.0
	pid.Err = 0.0
	pid.ErrLast = 0.0
	pid.Integral = 0.0
	pid.Kp = 0.31
	pid.Ki = 0.070
	pid.Kd = 0.3
}

func (pid *Pid) Realize(tempValue float64) float64 {
	pid.Err = pid.TargetValue - tempValue
	pid.Integral += pid.Err
	pid.ActualValue = pid.Kp*pid.Err + pid.Ki*pid.Integral + pid.Kd*(pid.Err-pid.ErrLast)
	pid.ErrLast = pid.Err
	return pid.ActualValue
}
