package xiaomi

import (
	"testing"
)

func TestSender_SendByRegid(t *testing.T) {
	sender := NewSender("peb+Sg9", "com.swxctx.app", true)
	regId := "rxyEIJOEWjKPRsx+FYxd4+FXyej67+A"
	_, err := sender.SendByRegid(&Message{
		Title:       "Swxctx",
		Description: "通知测试-regid",
		NotifyType:  -1,
		Target: Target{
			RegistrationId: regId,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
	}
}

func TestSender_SendByAlias(t *testing.T) {
	sender := NewSender("peb+Sg9", "com.swxctx.app", true)
	alias := "production-123"
	_, err := sender.SendByAlias(&Message{
		Title:       "Swxctx",
		Description: "通知测试-alias",
		NotifyType:  -1,
		Target: Target{
			Alias: alias,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
	}
}
