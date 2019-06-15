package WechatWork

import (
	"strconv"

	"github.com/dfang/wechat-work-go/models"
)

// resty "gopkg.in/resty.v1"

// 应用管理
// https://work.weixin.qq.com/api/doc?st=3F60C31C0B2943E1290B8A5AFE98479C0401174C41FCA19BF2E70120BB5A6D9A8E2A896089A560BA4FA3BDDAB38606AC2F558261EA0C2A42EC6B3E10B17CF9B768CC1FB27B6DF6F4AA28FD284C8368345A4F73D922387DF0BD5DD01710E8C00238AC37E06C8804735792C2283C485F4AABBBC322C446F5B03EF756AE083BE2C20A44979A45C41E58AFEA4D96388FF19F&vid=1688850695430171&cst=E706CCFA7EF5FF061CEB359C7C268C93AC42E97697E354FDB9234D01B77E85320E0A3EE1F3BDA68C59EBDCD6BE56AC01&deviceid=ef5b143f-cd0f-42bc-b0e4-507cb1b85737&version=2.7.8.2009&platform=mac#90000/90135/90227

// 创建成员
// 读取成员
// 更新成员
// 删除成员
// 批量删除成员
// 获取部门成员
// 获取部门成员详情

// GetMember 获取成员详情
func (c *App) GetMember(userID int64) (models.RespMemberGet, error) {
	req := models.ReqMemberGet{
		// AccessToken: c.AccessToken,
		UserID: userID,
	}

	// https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
	// resp, err := resty.R().Get("http://httpbin.org/get")

	var resp models.RespMemberGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		// TODO: error_chain
		return models.RespMemberGet{}, err
	}

	return resp, nil
}

// CreateMember 创建成员详情
func (c *App) CreateMember(userID int64) (models.RespMemberCreate, error) {
	req := models.ReqMemberCreate{
		UserID: strconv.FormatInt(userID, 10),
	}

	// https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
	// resp, err := resty.R().Get("http://httpbin.org/get")

	var resp models.RespMemberCreate
	err := c.executeQyapiJSONPost("/cgi-bin/user/create", req, &resp, true)
	if err != nil {
		// TODO: error_chain
		return models.RespMemberCreate{}, err
	}

	return resp, nil
}

// // SetAgent 设置应用详情
// func (c *App) SetAgent(req ReqAgentSet) (RespCommon, error) {
// 	// /cgi-bin/agent/set?access_token=ACCESS_TOKEN
// 	apiPath := "/cgi-bin/agent/set"
// 	var resp RespCommon
// 	err := c.executeQyapiJSONPost(apiPath, req, &resp, true)
// 	if err != nil {
// 		return RespCommon{}, err
// 	}

// 	return resp, nil
// }

// // ListAgent 获取应用列表
// func (c *App) ListAgent() error {
// 	return errors.New("not implemented")
// }

// func (c *App) createMenu() {
// 	// return errors.New("not implemented")
// }

// func (c *App) getMenu() error {
// 	// req := ReqAgentGet{
// 	// 	AccessToken: c.AccessToken,
// 	// 	AgentID:     c.AgentID,
// 	// }

// 	// var resp RespAgentGet
// 	// err := c.executeQyapiGet("/cgi-bin/menu/get", req, &resp, true)
// 	// if err != nil {
// 	// 	// TODO: error_chain
// 	// 	return RespAgentGet{}, err
// 	// }

// 	// return resp, nil
// 	return errors.New("not implemented")
// }

// func (c *App) deleteMenu() error {
// 	return errors.New("not implemented")
// }

// func SetAgent()
// func createMenu
// func getMenu
// func deleteMenu

// Member 成员
type Member struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	UserID         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Avatar         string `json:"avatar"`
	Telephone      string `json:"telephone"`
	Enable         int    `json:"enable"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		ExternalAttr     []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}
