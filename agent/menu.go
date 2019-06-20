package agent

import (
	"fmt"
)

// CreateMenu 创建菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
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
