package xiaomi

// doc: https://dev.mi.com/console/doc/detail?pId=230

// Sender
type Sender struct {
	// app secret
	Secret string
	// 包名，用逗号(,)分隔
	Packages string
	// 显示请求日志
	Debug bool
}

// NewSender
func NewSender(secret, packages string, debug ...bool) *Sender {
	sender := &Sender{
		Secret:   secret,
		Packages: packages,
	}
	if len(debug) > 0 {
		sender.Debug = debug[0]
	}
	return sender
}
