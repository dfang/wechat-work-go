// Package message 消息推送 API
//
// https://work.weixin.qq.com/api/doc#90000/90135/90235
package message

import (
	. "github.com/dfang/wechat-work-go/models"
)

// AppMessageCommon 公共字段
// touser，toparty, totag 不可同时为空
type AppMessageCommon struct {
	ToUser  string `json:"touser,omitempty"`
	ToParty string `json:"toparty,omitempty"`
	ToTag   string `json:"totag,omitempty"`

	AgentID int `json:"agentid"`

	MsgType string `json:"msgtype"`
	Safe    int    `json:"safe"`
}

// AppTextMessage 文字信息
type AppTextMessage struct {
	AppMessageCommon

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (a AppTextMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppImageMessage 图片信息
type AppImageMessage struct {
	AppMessageCommon

	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
}

func (a AppImageMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppVoiceMessage 语音消息
type AppVoiceMessage struct {
	AppMessageCommon

	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
}

func (a AppVoiceMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppVedioMessage 视频消息
type AppVedioMessage struct {
	AppMessageCommon

	Video struct {
		MediaID     string `json:"media_id"`
		Description string `json:"description"`
	} `json:"video"`
}

func (a AppVedioMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppFileMessage 文件消息
type AppFileMessage struct {
	AppMessageCommon

	File struct {
		MediaID string `json:"media_id"`
	} `json:"file"`
}

func (a AppFileMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppTextCardMessage 文本卡片消息
type AppTextCardMessage struct {
	AppMessageCommon

	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}

func (a AppTextCardMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppNewsMessage 图文消息
type AppNewsMessage struct {
	AppMessageCommon

	News struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
}

func (a AppNewsMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppNewsMessage 图文消息2
type AppMPNewsMessage struct {
	AppMessageCommon

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

func (a AppMPNewsMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

// AppMarkdownMessage markdown消息
type AppMarkdownMessage struct {
	AppMessageCommon

	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

func (a AppMarkdownMessage) Sendable() bool {
	return !(a.ToParty == "" && a.ToUser == "" && a.ToTag == "")
}

func NewAppTextMessage(agentID int, content string, safe int) AppTextMessage {
	a := AppTextMessage{}
	a.AgentID = agentID
	a.MsgType = "text"
	a.Text = struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	a.Safe = safe
	return a
}

func NewAppImageMessage(agentID int, mediaid string, safe int) AppImageMessage {
	a := AppImageMessage{}
	a.AgentID = agentID
	a.MsgType = "image"
	a.Image = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	a.Safe = safe
	return a
}

func NewAppVoiceMessage(agentID int, mediaid string) AppVoiceMessage {
	a := AppVoiceMessage{}
	a.AgentID = agentID
	a.MsgType = "voice"
	a.Voice = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	return a
}

func NewAppVedioMessage(agentID int, mediaid, description string, safe int) AppVedioMessage {
	a := AppVedioMessage{}
	a.AgentID = agentID
	a.MsgType = "vedio"
	a.Video = struct {
		MediaID     string `json:"media_id"`
		Description string `json:"description"`
	}{
		MediaID:     mediaid,
		Description: description,
	}
	a.Safe = safe
	return a
}

func NewAppFileMessage(agentID int, mediaid string, safe int) AppFileMessage {
	a := AppFileMessage{}
	a.AgentID = agentID
	a.MsgType = "file"
	a.File = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	a.Safe = safe
	return a
}

// RespSendAppMsg 发送应用消息 返回结果
type RespSendAppMsg struct {
	RespCommon

	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}
