package models

import (
	"encoding/json"
)

// RespAgentGet 查询应用响应
type RespAgentGet struct {
	RespCommon

	AgentID        int    `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoURL  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserinfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		Partyid []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		Tagid []int `json:"tagid"`
	} `json:"allow_tags"`
	Close              int    `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int    `json:"report_location_flag"`
	Isreportenter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

// ReqAgentSet 设置应用请求
type ReqAgentSet struct {
	AgentID            string `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag,omitempty"`
	LogoMediaID        string `json:"logo_mediaid,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	RedirectDomain     string `json:"redirect_domain,omitempty"`
	IsReportEnter      int    `json:"isreportenter,omitempty"`
	HomeURL            string `json:"home_url,omitempty"`
}

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for ReqAgentSet
func (x ReqAgentSet) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type RespAgentList struct {
	*RespCommon
	AgentList []AgentItem `json:"agentlist"`
}

type AgentItem struct {
	Agentid       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url"`
}
