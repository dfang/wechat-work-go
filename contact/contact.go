// Package contact 通讯录管理
//
// https://work.weixin.qq.com/api/doc#90000/90135/90193
package contact

import (
	wechatwork "github.com/dfang/wechat-work-go"
	. "github.com/dfang/wechat-work-go/models"
	"github.com/pkg/errors"
)

// Contact
type Contact struct {
	App *wechatwork.App
}

// CreateMember 创建成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90195
func (contact Contact) CreateMember(req ReqMemberCreate) (RespMemberCreate, error) {
	// resp, err := resty.R().Get("http://httpbin.org/get")
	apiPath := "cgi-bin/user/create"
	var data RespMemberCreate
	err := contact.App.Post(apiPath, nil, req, &data, true)
	if err != nil {
		return RespMemberCreate{}, err
	}
	return data, nil
}

// GetMember 获取成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90196
func (contact Contact) GetMember(userID string) (RespMemberGet, error) {
	apiPath := "/cgi-bin/user/get"
	qs := ReqMemberGet{
		UserID: userID,
	}
	var data RespMemberGet
	err := contact.App.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespMemberGet{}, err
	}
	return data, nil
}

// DeleteMember 删除成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90198
func (contact Contact) DeleteMember(userID string) (RespCommon, error) {
	apiPath := "/cgi-bin/user/delete"
	qs := ReqMemberGet{
		UserID: userID,
	}
	var data RespCommon
	err := contact.App.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// DeleteMembers 批量删除成员
//
// https://work.weixin.qq.com/api/doc#90000/90135/90199
func (contact Contact) DeleteMembers(userID string) error {
	return errors.New("not implemented")
}

// UpdateMember 更新成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90197
func (contact Contact) UpdateMember(body Member) (RespCommon, error) {
	apiPath := "/cgi-bin/user/update"
	qs := ReqMemberGet{
		AccessToken: contact.App.AccessToken,
	}
	var data RespCommon
	err := contact.App.Post(apiPath, qs, body, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// UserIDToOpenID userid转openid
//
// https://work.weixin.qq.com/api/doc#90000/90135/90202
func (contact Contact) UserIDToOpenID() error {
	return errors.New("not implemented")
}

// OpenIDToUserID userid转openid
//
// https://work.weixin.qq.com/api/doc#90000/90135/90202
func (app Contact) OpenIDToUserID() error {
	return errors.New("not implemented")
}

// TwoFactorAuth 二次验证
//
// https://work.weixin.qq.com/api/doc#90000/90135/90203
func (app Contact) TwoFactorAuth() error {
	return errors.New("not implemented")
}

// InviteMember 邀请成员
//
// https://work.weixin.qq.com/api/doc#90000/90135/90975
func (app Contact) InviteMember() error {
	return errors.New("not implemented")
}
