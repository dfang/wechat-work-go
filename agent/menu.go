package agent

import (
	"fmt"
)

// CreateMenu 创建菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
//
// note: 这里m的参数类型, 意味着你传入传入参数的时候既可以是Menu 也是其他自定义对象，甚至一段json字段都可以，非常的方便，因为resty会自动marshal, 只要符合企业微信菜单的结构就行了，具体见上述链接
func (agent *Agent) CreateMenu(agentID int64, m interface{}) (RespCommon, error) {
	apiPath := "/cgi-bin/menu/create"
	uri := agent.formatURL(apiPath, agentID)
	var result RespCommon
	err := agent.App.SimplePost(uri, m, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// GetMenu 获取菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90232
func (agent *Agent) GetMenu(agentID int64) (Menu, error) {
	apiPath := "/cgi-bin/menu/get"
	uri := agent.formatURL(apiPath, agentID)
	var result Menu
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return Menu{}, err
	}
	return result, nil
}

// DeleteMenu 删除菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90233
func (agent *Agent) DeleteMenu(agentID int64) (RespCommon, error) {
	apiPath := "/cgi-bin/menu/delete"
	uri := agent.formatURL(apiPath, agentID)
	var result RespCommon
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return RespCommon{}, nil
	}
	return result, nil
}

// formatUrl format url for Menu API
func (agent *Agent) formatURL(apiPath string, agentID int64) string {
	token := agent.App.GetAccessToken()
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%d", apiPath, token, agentID)
	return uri
}

// Menu 菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
type Menu struct {
	Button []Button `json:"button"`
}

// Button 菜单项
type Button struct {
	Name      string      `json:"name"`
	SubButton []SubButton `json:"sub_button,omitempty"`
	Type      string      `json:"type,omitempty"`
	Key       string      `json:"key,omitempty"`
}

// SubButton 子菜单项
type SubButton struct {
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	Key       string        `json:"key,omitempty"`
	URL       string        `json:"url,omitempty"`
	SubButton []interface{} `json:"sub_button,omitempty"`
}
