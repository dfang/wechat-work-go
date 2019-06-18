package contact

import (
	. "github.com/dfang/wechat-work-go/models"
	"github.com/pkg/errors"
)

// ListMembers 获取部门下成员概要
//
// https://work.weixin.qq.com/api/doc#90000/90135/90200
func (contact Contact) ListMembers(departmentID string, fetchChild bool) (RespListMembers, error) {
	apiPath := "/cgi-bin/user/simplelist"
	qs := ReqListMembers{
		DepartmentID: departmentID,
		FetchChild:   fetchChild,
	}
	var data RespListMembers
	err := contact.App.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespListMembers{}, err
	}
	return data, nil
}

// ListMembers2 获取部门下成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90201
func (contact Contact) ListMembers2(departmentID string, fetchChild bool) error {
	// apiPath := "/cgi-bin/user/list"
	return errors.New("not implemented")
}

// CreateDepartment 创建部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90205
func (contact Contact) CreateDepartment(req ReqCreateDepartment) (RespCreateDepartment, error) {
	apiPath := "cgi-bin/department/create"
	var data RespCreateDepartment
	err := contact.App.Post(apiPath, nil, req, &data, true)
	if err != nil {
		return RespCreateDepartment{}, err
	}
	return data, nil
}

// UpdateDepartment 更新部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90206
func (contact Contact) UpdateDepartment(body ReqUpdateDepartment) (RespCommon, error) {
	apiPath := "/cgi-bin/department/update"
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

// DeleteDepartment 删除部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90207
func (contact Contact) DeleteDepartment(departmentID string) (RespCommon, error) {
	apiPath := "/cgi-bin/department/delete"
	qs := ReqDepartmentDelete{
		DepartmentID: departmentID,
	}
	var data RespCommon
	err := contact.App.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// ListDepartments 获取部门列表
//
// https://work.weixin.qq.com/api/doc#90000/90135/90208
func (contact Contact) ListDepartments(departmentID string) (RespListDepartments, error) {
	apiPath := "/cgi-bin/department/list"
	qs := ReqListDepartments{
		DepartmentID: departmentID,
	}
	var data RespListDepartments
	err := contact.App.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespListDepartments{}, err
	}
	return data, nil
}
