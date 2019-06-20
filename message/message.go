// 应用推送消息
//
// https://work.weixin.qq.com/api/doc#90000/90135/90248

package message

import (
	"fmt"

	"github.com/dfang/wechat-work-go/models"
)

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
func (m Message) SendAppChatMessage(v Sendable) (models.RespCommon, error) {
	apiPath := "/cgi-bin/appchat/send"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())

	var result models.RespCommon
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return models.RespCommon{}, err
	}
	return result, nil
}
