package message

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"
)

// 群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90245

type GroupChat struct {
	App *wechatwork.App
}

// CreateGroupChat 创建群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90245
func (g GroupChat) CreateGroupChat(req ReqCreateGroupChat) (RespCreateGroupChat, error) {
	apiPath := "cgi-bin/appchat/create"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, g.App.GetAccessToken())
	var result RespCreateGroupChat
	err := g.App.SimplePost(uri, req, &result)
	if err != nil {
		return RespCreateGroupChat{}, err
	}
	return result, nil
}

// UpdateGroupChat 修改群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90246
func (g GroupChat) UpdateGroupChat(req ReqUpdateGroupChat) (models.RespCommon, error) {
	apiPath := "cgi-bin/appchat/update"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, g.App.GetAccessToken())
	var result models.RespCommon
	err := g.App.SimplePost(uri, req, &result)
	if err != nil {
		return models.RespCommon{}, err
	}
	return result, nil
}

// GetGroupChat 获取群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90247
func (g GroupChat) GetGroupChat(chatid string) (RespGetGroupChat, error) {
	apiPath := "cgi-bin/appchat/get"
	uri := fmt.Sprintf("%s?access_token=%s&chatid=%s", apiPath, g.App.GetAccessToken(), chatid)
	var result RespGetGroupChat
	err := g.App.SimpleGet(uri, &result)
	if err != nil {
		return RespGetGroupChat{}, err
	}
	return result, nil
}

type ReqCreateGroupChat struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatID   string   `json:"chatid"`
}

type RespCreateGroupChat struct {
	models.RespCommon

	UserList []string `json:"userlist"`
	ChatID   string   `json:"chatid"`
}

type ReqUpdateGroupChat struct {
	ChatID      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

type RespGetGroupChat struct {
	models.RespCommon

	ChatInfo struct {
		ChatID   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		UserList []string `json:"userlist"`
	} `json:"chat_info"`
}
