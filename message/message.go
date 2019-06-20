// 应用推送消息
//
// https://work.weixin.qq.com/api/doc#90000/90135/90248

package message

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
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

type Sendable interface {
	Sendable() bool
}

// SendAppMessage 发送应用消息 (发送到应用的)
//
// https://work.weixin.qq.com/api/doc#90000/90135/90236
func (m Message) SendAppMessage(v Sendable) (RespSendAppMsg, error) {
	apiPath := "/cgi-bin/message/send"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())

	if !v.Sendable() {
		panic("touser, toparty, totag 不能同时为空, agentid 不能为空, 请检查")
	}

	var result RespSendAppMsg
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return RespSendAppMsg{}, err
	}
	return result, nil
}

// SendGroupChatMessage 应用推送消息 (发送到群聊的)
//
// https://work.weixin.qq.com/api/doc#90000/90135/90248
func (m Message) SendGroupChatMessage(v Sendable) (RespCommon, error) {
	apiPath := "/cgi-bin/appchat/send"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())

	var result RespCommon
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// UpdateTaskCard 更新任务卡片消息状态
//
// https://work.weixin.qq.com/api/doc#90000/90135/91579
func (m Message) UpdateTaskCard(v ReqTaskCardUpdate) (RespTaskCardUpdate, error) {
	apiPath := "/cgi-bin/message/update_taskcard"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())

	var result RespTaskCardUpdate
	err := m.App.SimplePost(uri, v, &result)
	if err != nil {
		return RespTaskCardUpdate{}, err
	}
	return result, nil
}

// GetCallbackIPs 获取企业微信服务器的ip段
//
// https://work.weixin.qq.com/api/doc#90000/90135/90237/%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%BE%AE%E4%BF%A1%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%9A%84ip%E6%AE%B5
func (m Message) GetCallbackIPs(v ReqTaskCardUpdate) (RespGetCallbackIPs, error) {
	apiPath := "/cgi-bin/getcallbackip"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, m.App.GetAccessToken())
	var result RespGetCallbackIPs
	err := m.App.SimpleGet(uri, &result)
	if err != nil {
		return RespGetCallbackIPs{}, err
	}
	return result, nil
}

// ReqTaskCardUpdate 更新任务卡片状态的请求
type ReqTaskCardUpdate struct {
	Userids    []string `json:"userids"`
	Agentid    int      `json:"agentid"`
	TaskID     string   `json:"task_id"`
	ClickedKey string   `json:"clicked_key"`
}

// RespTaskCardUpdate 更新任务卡片状态的响应
type RespTaskCardUpdate struct {
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
	Invaliduser []string `json:"invaliduser"`
}

type RespGetCallbackIPs struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	IPList  []string `json:"ip_list"`
}
