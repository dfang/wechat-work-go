package message

import (
	"fmt"

	"github.com/dfang/wechat-work-go/models"
)

// 群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90245

// type GroupChat struct {
// 	App *wechatwork.App
// }

// CreateGroupChat 创建群聊会话
//
// https://work.weixin.qq.com/api/doc#90000/90135/90245
func (g Message) CreateGroupChat(req ReqCreateGroupChat) (RespCreateGroupChat, error) {
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
func (g Message) UpdateGroupChat(req ReqUpdateGroupChat) (models.RespCommon, error) {
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
func (g Message) GetGroupChat(chatid string) (RespGetGroupChat, error) {
	apiPath := "cgi-bin/appchat/get"
	uri := fmt.Sprintf("%s?access_token=%s&chatid=%s", apiPath, g.App.GetAccessToken(), chatid)
	var result RespGetGroupChat
	err := g.App.SimpleGet(uri, &result)
	if err != nil {
		return RespGetGroupChat{}, err
	}
	return result, nil
}

// ReqCreateGroupChat 创建群聊会话请求
type ReqCreateGroupChat struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatID   string   `json:"chatid"`
}

// RespCreateGroupChat 创建群聊会话响应
type RespCreateGroupChat struct {
	models.RespCommon

	ChatID string `json:"chatid"`
}

// ReqUpdateGroupChat 修改群聊会话请求
type ReqUpdateGroupChat struct {
	ChatID      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

// RespGetGroupChat 获取群聊会话响应
type RespGetGroupChat struct {
	models.RespCommon

	ChatInfo struct {
		ChatID   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		UserList []string `json:"userlist"`
	} `json:"chat_info"`
}

// GroupChatMessageCommon 公共字段
type GroupChatMessageCommon struct {
	ChatID string `json:"chatid"`

	MsgType string `json:"msgtype"`
	Safe    int    `json:"safe"`
}

// AppTextMessage 文字信息
type GroupChatTextMessage struct {
	GroupChatMessageCommon

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (a GroupChatTextMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppImageMessage 图片信息
type GroupChatImageMessage struct {
	GroupChatMessageCommon

	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
}

func (a GroupChatImageMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppVoiceMessage 语音消息
type GroupChatVoiceMessage struct {
	GroupChatMessageCommon

	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
}

func (a GroupChatVoiceMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppVedioMessage 视频消息
type GroupChatVedioMessage struct {
	GroupChatMessageCommon

	Video struct {
		MediaID     string `json:"media_id"`
		Description string `json:"description"`
	} `json:"video"`
}

func (a GroupChatVedioMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppFileMessage 文件消息
type GroupChatFileMessage struct {
	GroupChatMessageCommon

	File struct {
		MediaID string `json:"media_id"`
	} `json:"file"`
}

func (a GroupChatFileMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppTextCardMessage 文本卡片消息
type GroupChatTextCardMessage struct {
	GroupChatMessageCommon

	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}

func (a GroupChatTextCardMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppNewsMessage 图文消息
type GroupChatNewsMessage struct {
	GroupChatMessageCommon

	News struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
}

func (a GroupChatNewsMessage) Sendable() bool {
	return a.ChatID != ""
}

// AppNewsMessage 图文消息2
type GroupChatMPNewsMessage struct {
	GroupChatMessageCommon

	MPNews struct {
		Articles []struct {
			Title            string `json:"title"`
			ThumbMediaID     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceURL string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		} `json:"articles"`
	} `json:"mpnews"`
	Safe int `json:"safe"`
}

func (a GroupChatMPNewsMessage) Sendable() bool {
	return a.ChatID != ""
}
