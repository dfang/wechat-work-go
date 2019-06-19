package agent

import (
	"fmt"

	"github.com/dfang/wechat-work-go/models"
)

// CreateMenu 创建菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
func (agent *Agent) CreateMenu(agentID string, m models.Menu) (models.RespCommon, error) {
	apiPath := "/cgi-bin/menu/create"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, agent.App.GetAccessToken(), agentID)
	var result models.RespCommon
	err := agent.App.SimplePost(uri, m, &result)
	if err != nil {
		return models.RespCommon{}, err
	}
	return result, nil
}

// GetMenu 获取菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90232
func (agent *Agent) GetMenu(agentID string) (models.Menu, error) {
	apiPath := "/cgi-bin/menu/get"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, agent.App.GetAccessToken(), agentID)
	var result models.Menu
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return models.Menu{}, err
	}
	return result, nil
}

// DeleteMenu 删除菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90233
func (agent *Agent) DeleteMenu(agentID string) (models.RespCommon, error) {
	apiPath := "/cgi-bin/menu/delete"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, agent.App.GetAccessToken(), agentID)
	var result models.RespCommon
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return models.RespCommon{}, nil
	}
	return result, nil
}
