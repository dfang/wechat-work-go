// 应用推送消息
//
// https://work.weixin.qq.com/api/doc#90000/90135/90248

package message

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
	. "github.com/dfang/wechat-work-go/models"
)

// Message 消息推送
type Message struct {
	App *wechatwork.App
}

// WithApp 返回 Message 的实例
//
// 所有消息推送相关API 通过此方法返回的实例调用
func WithApp(app *wechatwork.App) *Message {
	return &Message{
		App: app,
	}
}

type Sendable interface{}

// SendMessage 发送应用消息
//
// https://work.weixin.qq.com/api/doc#90000/90135/90236
func (m Message) SendMessage(v Sendable) (RespSendMsg, error) {
	apiPath := "/cgi-bin/message/send"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())
	var result RespSendMsg
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return RespSendMsg{}, err
	}
	return result, nil
}

// SendAppChatMessage 应用推送消息
//
// https://work.weixin.qq.com/api/doc#90000/90135/90248
func (m Message) SendAppChatMessage(v Sendable) (RespCommon, error) {
	apiPath := "/cgi-bin/appchat/send"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())

	var result RespCommon
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}
