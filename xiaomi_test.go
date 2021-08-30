package push

import (
	"testing"

	"github.com/swxctx/push/xiaomi"
)

func Test_XiaomiSendByRegid(t *testing.T) {
	sender := xiaomi.NewSender("peb+Sg9", "com.swxctx.app", true)
	regId := "rxyEIJOEWjKPRsx+FYxd4+FXyej67+A"
	_, err := sender.SendByRegid(&xiaomi.Message{
		Title:       "Swxctx",
		Description: "通知测试-regid",
		NotifyType:  -1,
		Target: xiaomi.Target{
			RegistrationId: regId,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
	}
}

func Test_XiaomiSendByAlias(t *testing.T) {
	sender := xiaomi.NewSender("peb+Sg9", "com.swxctx.app", true)
	message := &xiaomi.Message{
		Title:       "Swxctx",
		Description: "通知测试-regid",
		NotifyType:  -1,
	}
	message.SetAlias("12345")
	_, err := sender.SendByRegid(message)
	if err != nil {
		t.Errorf("err-> %v", err)
	}
}
