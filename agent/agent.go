// Package agent 提供应用管理相关的接口（除创建）
//
// 注意: 目前不支持通过API接口创建应用, 需要去企业微信管理后台手动创建
//
// API 文档链接: https://work.weixin.qq.com/api/doc#90000/90135/90226
package agent

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"
)

// Agent 应用管理
type Agent struct {
	App *wechatwork.App
}

// WithApp 返回 Agent 实例
//
// 所有应用管理相关API 通过此方法返回的实例调用
func WithApp(app *wechatwork.App) *Agent {
	return &Agent{
		App: app,
	}
}

// GetAgent 获取应用
//
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (agent *Agent) GetAgent(agentID string) (models.RespAgentGet, error) {
	apiPath := "/cgi-bin/agent/get"
	uri := fmt.Sprintf("%s?access_token=%s&agentid=%s", apiPath, agent.App.GetAccessToken(), agentID)
	var result models.RespAgentGet
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return models.RespAgentGet{}, err
	}
	return result, nil
}

// ListAgents 获取access_token 下应用列表
//
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (agent *Agent) ListAgents() (models.RespAgentList, error) {
	apiPath := "/cgi-bin/agent/list"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, agent.App.GetAccessToken())
	var result models.RespAgentList
	err := agent.App.SimpleGet(uri, &result)
	if err != nil {
		return models.RespAgentList{}, err
	}
	return result, nil
}

// SetAgent 设置应用详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90228
func (agent *Agent) SetAgent(req models.ReqAgentSet) (models.RespCommon, error) {
	apiPath := "/cgi-bin/agent/set"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, agent.App.GetAccessToken())
	var result models.RespCommon
	err := agent.App.SimplePost(uri, req, &result)
	if err != nil {
		return models.RespCommon{}, err
	}
	return result, nil
}
