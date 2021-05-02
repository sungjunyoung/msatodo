package alarm

type Alarm interface {
	Register()
	Send()
}
