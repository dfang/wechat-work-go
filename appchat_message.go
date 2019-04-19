package WechatWork

import "encoding/json"

// AppChatMessage 应用消息
type AppChatMessage struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
}

// AppChatTextMessage 文字信息
type AppChatTextMessage struct {
	AppChatMessage

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

// AppChatImageMessage 图片信息
type AppChatImageMessage struct {
	AppChatMessage

	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
	Safe int `json:"safe"`
}

// AppChatVoiceMessage 语音消息
type AppChatVoiceMessage struct {
	AppChatMessage

	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
}

// AppChatVedioMessage 视频消息
type AppChatVedioMessage struct {
	AppChatMessage

	Video struct {
		MediaID     string `json:"media_id"`
		Description string `json:"description"`
	} `json:"video"`
	Safe int `json:"safe"`
}

// AppChatFileMessage 文件消息
type AppChatFileMessage struct {
	AppChatMessage

	File struct {
		MediaID string `json:"media_id"`
	} `json:"file"`
	Safe int `json:"safe"`
}

// AppChatTextCardMessage 文本卡片消息
type AppChatTextCardMessage struct {
	AppChatMessage

	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
	Safe int `json:"safe"`
}

// AppChatNewsMessage 图文消息
type AppChatNewsMessage struct {
	AppChatMessage

	News struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
	Safe int `json:"safe"`
}

// AppChatMarkdownMessage markdown消息
type AppChatMarkdownMessage struct {
	AppChatMessage

	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

func NewAppChatTextMessage(chatid, content string, safe int) AppChatTextMessage {
	a := AppChatTextMessage{}
	a.Chatid = chatid
	a.Msgtype = "text"
	a.Text = struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	a.Safe = safe
	return a
}

func NewAppChatImageMessage(chatid, mediaid string, safe int) AppChatImageMessage {
	a := AppChatImageMessage{}
	a.Chatid = chatid
	a.Msgtype = "image"
	a.Image = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	a.Safe = safe
	return a
}

func NewAppChatVoiceMessage(chatid, mediaid string) AppChatVoiceMessage {
	a := AppChatVoiceMessage{}
	a.Chatid = chatid
	a.Msgtype = "voice"
	a.Voice = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	return a
}

func NewAppChatVedioMessage(chatid, mediaid, description string, safe int) AppChatVedioMessage {
	a := AppChatVedioMessage{}
	a.Chatid = chatid
	a.Msgtype = "vedio"
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

func NewAppChatFileMessage(chatid, mediaid string, safe int) AppChatFileMessage {
	a := AppChatFileMessage{}
	a.Chatid = chatid
	a.Msgtype = "file"
	a.File = struct {
		MediaID string `json:"media_id"`
	}{
		MediaID: mediaid,
	}
	a.Safe = safe
	return a
}

// func NewAppChatNewsMessage(chatid, mediaid string, safe int) AppChatNewsMessage {
// 	a := AppChatNewsMessage{}
// 	a.Chatid = chatid
// 	a.Msgtype = "news"
// 	a.News = struct {
// 			Articles []struct {
// 				Title       string `json:"title"`
// 				Description string `json:"description"`
// 				URL         string `json:"url"`
// 				Picurl      string `json:"picurl"`
// 			} `json:"articles"`
// 		} {
// 			Articles: []struct {

// 			},
// 		}
// a.Safe = safe
// 	return a
// }

func NewAppChatTextCardMessage(chatid, title, description, url, btntxt string, safe int) AppChatTextCardMessage {
	a := AppChatTextCardMessage{}
	a.Chatid = chatid
	a.Msgtype = "textcard"
	a.Textcard = struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	}{
		Title:       title,
		Description: description,
		URL:         url,
		Btntxt:      btntxt,
	}
	a.Safe = safe
	return a
}

func NewAppChatMarkdownMessage(chatid, content string) AppChatMarkdownMessage {
	a := AppChatMarkdownMessage{}
	a.Chatid = chatid
	a.Msgtype = "markdown"
	a.Markdown = struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	return a
}

func (x AppChatTextMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatImageMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatVoiceMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatVedioMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatFileMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatTextCardMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatNewsMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (x AppChatMarkdownMessage) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}
