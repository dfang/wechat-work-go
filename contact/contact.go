// Package contact 提供通讯录管理相关的接口
//
// 注意: 关于创建成员（客服答复）
//
// 目前只能使用通讯录的secret 获取token进行创建  其他的secret是没有创建成员的权限的
//
// 获取路径：通讯录管理secret。在“管理工具”-“通讯录同步”里面查看（需开启“API接口同步”）
//
// https://work.weixin.qq.com/api/doc#90000/90135/90193
package contact

import (
	"fmt"

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
	apiPath := "cgi-bin/user/create"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespMemberCreate
	err := contact.App.SimplePost(uri, req, &result)
	if err != nil {
		return RespMemberCreate{}, err
	}
	return result, nil
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
