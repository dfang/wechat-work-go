package wechatwork

import (
	. "github.com/dfang/wechat-work-go/models"
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

// CreateMember 创建成员详情
//
// https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN
func (app *App) CreateMember(req ReqMemberCreate) (RespMemberCreate, error) {
	// resp, err := resty.R().Get("http://httpbin.org/get")
	apiPath := "cgi-bin/user/create"
	var data RespMemberCreate
	err := app.Post(apiPath, nil, req, &data, true)
	if err != nil {
		return RespMemberCreate{}, err
	}
	return data, nil
}

// GetMember 获取成员详情
//
// https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
func (app *App) GetMember(userID string) (RespMemberGet, error) {
	apiPath := "/cgi-bin/user/get"
	qs := ReqMemberGet{
		UserID: userID,
	}
	var data RespMemberGet
	err := app.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespMemberGet{}, err
	}
	return data, nil
}

// DeleteMember 删除成员详情
//
//https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=USERID
func (app *App) DeleteMember(userID string) (RespCommon, error) {
	apiPath := "/cgi-bin/user/delete"
	qs := ReqMemberGet{
		UserID: userID,
	}
	var data RespCommon
	err := app.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// UpdateMember 更新成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90197
func (app *App) UpdateMember(body Member) (RespCommon, error) {
	apiPath := "/cgi-bin/user/update"
	qs := ReqMemberGet{
		AccessToken: app.AccessToken,
	}
	var data RespCommon
	err := app.Post(apiPath, qs, body, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// ListMembers 获取部门下成员
//
// https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
func (app *App) ListMembers(departmentID string, fetch_child bool) (RespListMembers, error) {
	apiPath := "/cgi-bin/user/simplelist"
	qs := ReqListMembers{
		DepartmentID: departmentID,
		FetchChild:   fetch_child,
	}
	var data RespListMembers
	err := app.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespListMembers{}, err
	}
	return data, nil
}

// CreateDepartment 创建部门
//
// https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN
func (app *App) CreateDepartment(req ReqCreateDepartment) (RespCreateDepartment, error) {
	apiPath := "cgi-bin/department/create"
	var data RespCreateDepartment
	err := app.Post(apiPath, nil, req, &data, true)
	if err != nil {
		return RespCreateDepartment{}, err
	}
	return data, nil
}

// UpdateDepartment 更新部门
//
// https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN
func (app *App) UpdateDepartment(body ReqUpdateDepartment) (RespCommon, error) {
	apiPath := "/cgi-bin/department/update"
	qs := ReqMemberGet{
		AccessToken: app.AccessToken,
	}
	var data RespCommon
	err := app.Post(apiPath, qs, body, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// DeleteDepartment 删除部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90207
//
// https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=ID
func (app *App) DeleteDepartment(departmentID string) (RespCommon, error) {
	apiPath := "/cgi-bin/department/delete"
	qs := ReqDepartmentDelete{
		DepartmentID: departmentID,
	}
	var data RespCommon
	err := app.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespCommon{}, err
	}
	return data, nil
}

// ListDepartments 获取部门列表
//
// https://work.weixin.qq.com/api/doc#90000/90135/90208
//
// https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID
func (app *App) ListDepartments(departmentID string) (RespListDepartments, error) {
	apiPath := "/cgi-bin/department/list"
	qs := ReqListDepartments{
		DepartmentID: departmentID,
	}
	var data RespListDepartments
	err := app.Get(apiPath, qs, &data, true)
	if err != nil {
		return RespListDepartments{}, err
	}
	return data, nil
}
