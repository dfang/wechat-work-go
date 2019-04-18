package WechatWork

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type reqAccessToken struct {
	CorpID     string
	CorpSecret string
}

// IntoURLValues 转换为 url.Values 类型
//
// impl urlValuer for reqAccessToken
func (x reqAccessToken) IntoURLValues() url.Values {
	return url.Values{
		"corpid":     {x.CorpID},
		"corpsecret": {x.CorpSecret},
	}
}

type RespCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *RespCommon) IsOK() bool {
	return x.ErrCode == 0
}

type respAccessToken struct {
	RespCommon

	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}

// reqMessage 消息发送请求
type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for reqMessage
func (x reqMessage) IntoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	result, err := json.Marshal(obj)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// respMessageSend 消息发送响应
type respMessageSend struct {
	RespCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}

type reqUserGet struct {
	UserID string
}

// IntoURLValues 转换为 url.Values 类型
//
// impl urlValuer for reqUserGet
func (x reqUserGet) IntoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// respUserGet 读取成员响应
type respUserGet struct {
	RespCommon

	UserID         string   `json:"userid"`
	Name           string   `json:"name"`
	DeptIDs        []int64  `json:"department"`
	DeptOrder      []uint32 `json:"order"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	AvatarURL      string   `json:"avatar"`
	Telephone      string   `json:"telephone"`
	IsEnabled      int      `json:"enable"`
	Alias          string   `json:"alias"`
	Status         int      `json:"status"`
	QRCodeURL      string   `json:"qr_code"`
	// TODO: extattr external_profile external_position
}

// reqAgentGet 查询应用请求
type reqAgentGet struct {
	AgentID     int64
	AccessToken string `json:"access_token"`
}

// IntoURLValues 转换为 url.Values 类型
//
// impl urlValuer for reqAgentGet
func (x reqAgentGet) IntoURLValues() url.Values {
	return url.Values{
		"agentid":      {strconv.FormatInt(x.AgentID, 10)},
		"access_token": {x.AccessToken},
	}
}

// RespAgentGet 查询应用响应
type RespAgentGet struct {
	RespCommon

	Agentid        int    `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoURL  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserinfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		Partyid []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		Tagid []int `json:"tagid"`
	} `json:"allow_tags"`
	Close              int    `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int    `json:"report_location_flag"`
	Isreportenter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

// ReqAgentSet 设置应用请求
type ReqAgentSet struct {
	Agentid            int    `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag,omitempty"`
	LogoMediaid        string `json:"logo_mediaid,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	RedirectDomain     string `json:"redirect_domain,omitempty"`
	Isreportenter      int    `json:"isreportenter,omitempty"`
	HomeURL            string `json:"home_url,omitempty"`
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for reqMessage
func (x ReqAgentSet) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 创建自定义菜单
type ReqMenuCreate struct {
	Button []struct {
		Name      string `json:"name"`
		Type      string `json:"type,omitempty"`
		Key       string `json:"key,omitempty"`
		SubButton []struct {
			Type      string        `json:"type"`
			Name      string        `json:"name"`
			Key       string        `json:"key"`
			URL       string        `json:"url,omitempty"`
			SubButton []interface{} `json:"sub_button"`
		} `json:"sub_button,omitempty"`
	} `json:"button"`
}
