package WechatWork

import (
	"errors"

	"github.com/dfang/wechat-work-go/models"
)

// 应用管理
// https://work.weixin.qq.com/api/doc?st=3F60C31C0B2943E1290B8A5AFE98479C0401174C41FCA19BF2E70120BB5A6D9A8E2A896089A560BA4FA3BDDAB38606AC2F558261EA0C2A42EC6B3E10B17CF9B768CC1FB27B6DF6F4AA28FD284C8368345A4F73D922387DF0BD5DD01710E8C00238AC37E06C8804735792C2283C485F4AABBBC322C446F5B03EF756AE083BE2C20A44979A45C41E58AFEA4D96388FF19F&vid=1688850695430171&cst=E706CCFA7EF5FF061CEB359C7C268C93AC42E97697E354FDB9234D01B77E85320E0A3EE1F3BDA68C59EBDCD6BE56AC01&deviceid=ef5b143f-cd0f-42bc-b0e4-507cb1b85737&version=2.7.8.2009&platform=mac#90000/90135/90227

// 获取应用
// 设置应用
// 创建菜单
// 获取菜单
// 删除菜单

// GetAgent 获取应用详情
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (c *App) GetAgent(agentID string) (models.RespAgentGet, error) {
	apiPath := "/cgi-bin/agent/get"
	req := models.ReqAgentGet{
		AgentID: agentID,
	}
	var resp models.RespAgentGet
	err := c.Get(apiPath, req, resp, true)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// SetAgent 设置应用详情
// https://work.weixin.qq.com/api/doc#90000/90135/90228
func (c *App) SetAgent(req models.ReqAgentSet) (models.RespCommon, error) {
	// /cgi-bin/agent/set?access_token=ACCESS_TOKEN
	// apiPath := "/cgi-bin/agent/set"
	// resty.SetHostURL("https://qyapi.weixin.qq.com")
	// resty.R().SetBody(req)
	// resp, err := resty.R().Post(apiPath)

	var resp models.RespCommon
	// err := c.executeQyapiJSONPost(apiPath, req, &resp, true)
	// if err != nil {
	// 	return models.RespCommon{}, err
	// }

	return resp, nil
}

// ListAgents 获取应用列表
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (c *App) ListAgents() (models.RespAgentList, error) {
	apiPath := "/cgi-bin/agent/list"
	var resp models.RespAgentList
	err := c.Get(apiPath, nil, &resp, true)
	if err != nil {
		return models.RespAgentList{}, err
	}
	return resp, nil
}

func (c *App) createMenu() error {
	return errors.New("not implemented")
}

func (c *App) getMenu() error {
	// req := ReqAgentGet{
	// 	AccessToken: c.AccessToken,
	// 	AgentID:     c.AgentID,
	// }

	// var resp RespAgentGet
	// err := c.executeQyapiGet("/cgi-bin/menu/get", req, &resp, true)
	// if err != nil {
	// 	// TODO: error_chain
	// 	return RespAgentGet{}, err
	// }

	// return resp, nil
	return errors.New("not implemented")
}

func (c *App) deleteMenu() error {
	return errors.New("not implemented")
}
