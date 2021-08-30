package xiaomi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/swxctx/ghttp"
	"github.com/swxctx/query"
)

// SendByRegid 向某个regid或一组regid列表推送某条消息（这些regId可以属于不同的包名）
func (sender *Sender) SendByRegid(message *Message) (*SendResponse, error) {
	return sender.sendMessage(message, BaseReqUrl+"/v3/message/regid")
}

// SendByAlias 向某个alias或一组alias列表推送某条消息（这些alias可以属于不同的包名）
func (sender *Sender) SendByAlias(message *Message) (*SendResponse, error) {
	return sender.sendMessage(message, BaseReqUrl+"/v3/message/alias")
}

// SendByTopic 向某个topic推送某条消息（可以指定一个或多个包名）
func (sender *Sender) SendByTopic(message *Message) (*SendResponse, error) {
	return sender.sendMessage(message, BaseReqUrl+"/v3/message/topic")
}

// SendByAlias 向多个topic推送单条消息（可以指定一个或多个包名）
func (sender *Sender) SendByMultiTopic(message *Message) (*SendResponse, error) {
	return sender.sendMessage(message, BaseReqUrl+"/v3/message/multi_topic")
}

// SendByAlias 向所有设备推送某条消息（可以指定一个或多个包名）
func (sender *Sender) SendByAll(message *Message) (*SendResponse, error) {
	return sender.sendMessage(message, BaseReqUrl+"/v3/message/all")
}

func (sender *Sender) sendMessage(message *Message, api string) (*SendResponse, error) {
	// 推送包名
	message.RestrictedPackageName = sender.Packages

	var (
		err  error
		form url.Values
	)
	form, err = query.NewEncoder().Encode(message)
	if err != nil {
		return nil, err
	}

	// 请求结构
	req := ghttp.Request{
		Url:       api,
		Body:      form.Encode(),
		Method:    "POST",
		ShowDebug: sender.Debug,
	}
	req.AddHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.AddHeader("Authorization", "key="+sender.Secret)

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// read resp body
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var (
		response *SendResponse
	)
	if err := json.Unmarshal(respBs, &response); err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return response, fmt.Errorf("send msg err, reason-> %s", response.Reason)
	}
	return nil, nil
}
