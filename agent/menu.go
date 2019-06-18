package agent

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"
)

type Menu struct {
	App *wechatwork.App
}

// CreateMenu 创建菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90231
func (menu Menu) CreateMenu(agentID string, m models.Menu) (models.RespCommon, error) {
	apiPath := "/cgi-bin/menu/create"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, menu.App.GetAccessToken(), agentID)
	// qs := url.Values{}
	// qs.Add("agentid", agentID)
	var result models.RespCommon
	err := menu.App.SimplePost(uri, m, &result)
	if err != nil {
		return models.RespCommon{}, err
	}
	return result, nil
}

// GetMenu 获取菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90232
func (menu Menu) GetMenu(agentID string) (models.Menu, error) {
	apiPath := "/cgi-bin/menu/get"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, menu.App.GetAccessToken(), agentID)
	var result models.Menu
	err := menu.App.SimpleGet(uri, &result)
	if err != nil {
		return models.Menu{}, err
	}
	return result, nil
}

// DeleteMenu 删除菜单
//
// https://work.weixin.qq.com/api/doc#90000/90135/90233
func (menu Menu) DeleteMenu(agentID string) (models.RespCommon, error) {
	apiPath := "/cgi-bin/menu/delete"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, menu.App.GetAccessToken(), agentID)
	var result models.RespCommon
	err := menu.App.SimpleGet(uri, &result)
	if err != nil {
		return models.RespCommon{}, nil
	}
	return result, nil
}
