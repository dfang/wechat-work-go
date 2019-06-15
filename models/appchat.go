package models

import (
	"encoding/json"
	"net/url"
)

// ReqAppChatCreate 创建群聊会话请求
type ReqAppChatCreate struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid"`
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for ReqAppChatCreate
func (x ReqAppChatCreate) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespAppChatCreate 创建群聊会话响应
type RespAppChatCreate struct {
	RespCommon
	CHATID string `json:"chatid"`
}

// ReqAppChatUpdate 修改群聊会话请求
type ReqAppChatUpdate struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	AddUserList []string `json:"add_user_list,omitempty"`
	DelUserList []string `json:"del_user_list,omitempty"`
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for ReqAppChatUpdate
func (x ReqAppChatUpdate) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReqAppChatGet 查询应用请求
type ReqAppChatGet struct {
	ChatID      string
	AccessToken string `json:"access_token"`
}

// IntoURLValues 转换为 url.Values 类型
//
// impl urlValuer for ReqAppChatGet
func (x ReqAppChatGet) IntoURLValues() url.Values {
	return url.Values{
		"chatid":       {x.ChatID},
		"access_token": {x.AccessToken},
	}
}

// RespAppChatGet 获取群聊会话响应
type RespAppChatGet struct {
	RespCommon

	ChatInfo struct {
		Chatid   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		Userlist []string `json:"userlist"`
	} `json:"chat_info"`
}
