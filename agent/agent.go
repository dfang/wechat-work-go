// Package agent 提供应用管理相关的接口（除创建）
//
// 注意: 目前不支持通过API接口创建应用, 需要去企业微信管理后台手动创建
//
// API 文档链接: https://work.weixin.qq.com/api/doc#90000/90135/90226
package agent

import (
	wechatwork "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/models"
)

// Agent 应用管理
type Agent struct {
	App *wechatwork.App
}

// GetAgent 获取应用
//
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (agent Agent) GetAgent(agentID string) (models.RespAgentGet, error) {
	apiPath := "/cgi-bin/agent/get"
	req := models.ReqAgentGet{
		AgentID: agentID,
	}
	var data models.RespAgentGet
	err := agent.App.Get(apiPath, req, &data, true)
	if err != nil {
		return models.RespAgentGet{}, err
	}
	return data, nil
}

// ListAgents 获取access_token 下应用列表
//
// https://work.weixin.qq.com/api/doc#90000/90135/90227
func (agent Agent) ListAgents() (models.RespAgentList, error) {
	apiPath := "/cgi-bin/agent/list"
	var data models.RespAgentList
	err := agent.App.Get(apiPath, nil, &data, true)
	if err != nil {
		return models.RespAgentList{}, err
	}
	return data, nil
}

// SetAgent 设置应用详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90228
func (agent Agent) SetAgent(req models.ReqAgentSet) (models.RespCommon, error) {
	// /cgi-bin/agent/set?access_token=ACCESS_TOKEN
	apiPath := "/cgi-bin/agent/set"
	var data models.RespCommon
	err := agent.App.Post(apiPath, nil, req, &data, true)
	if err != nil {
		return models.RespCommon{}, err
	}

	return data, nil
}
