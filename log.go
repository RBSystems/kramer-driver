package kramer

type Logger interface {
	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

func (v *Via) infof(format string, a ...interface{}) {
	if v.Logger != nil {
		v.Logger.Infof(format, a...)
	}
}