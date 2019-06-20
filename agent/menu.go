package agent

import (
	"fmt"
)

// CreateMenu 创建菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
func (agent *Agent) CreateMenu(agentID string, m Menu) (RespCommon, error) {
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
func (agent *Agent) GetMenu(agentID string) (Menu, error) {
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
func (agent *Agent) DeleteMenu(agentID string) (RespCommon, error) {
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
func (agent *Agent) formatURL(apiPath, agentID string) string {
	token := agent.App.GetAccessToken()
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, token, agentID)
	return uri
}
