// Package agent 提供应用管理相关的接口（除创建）
//
// 注意: 目前不支持通过API接口创建应用, 需要去企业微信管理后台手动创建
//
// API 文档链接: https://work.weixin.qq.com/api/doc#90000/90135/90226
package agent

import wechatwork "github.com/dfang/wechat-work-go"

// Agent 应用管理
type Agent struct {
	App *wechatwork.App
}

// GetAgent
// UpdateAgent
