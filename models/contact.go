package models

import (
	"encoding/json"
	"net/url"
)

// ReqMemberGet 查询成员请求
type ReqMemberGet struct {
	UserID      string `json:"userid"`
	AccessToken string `json:"access_token"`
}

// IntoURLValues impl url.valuer
func (x ReqMemberGet) IntoURLValues() url.Values {
	return url.Values{
		"userid":       {x.UserID},
		"access_token": {x.AccessToken},
	}
}

// RespMemberGet 查询成员响应
type RespMemberGet struct {
	RespCommon

	Member
}

type ReqMemberCreate struct {
	UserID         string `json:"userid"`
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	Mobile         string `json:"mobile"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Enable         int    `json:"enable"`
	AvatarMediaid  string `json:"avatar_mediaid"`
	Telephone      string `json:"telephone"`
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
	ToInvite         bool   `json:"to_invite"`
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

// IntoBody 转换为请求体的 []byte 类型
//
// impl bodyer for ReqMemberCreate
func (x ReqMemberCreate) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

type RespMemberCreate struct {
	RespCommon
}

// Member 成员
type Member struct {
	UserID           string `json:"userid"`
	Name             string `json:"name"`
	Department       []int  `json:"department"`
	Order            []int  `json:"order"`
	Position         string `json:"position"`
	Mobile           string `json:"mobile"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	IsLeaderInDept   []int  `json:"is_leader_in_dept"`
	Avatar           string `json:"avatar"`
	Telephone        string `json:"telephone"`
	Enable           int    `json:"enable"`
	Alias            string `json:"alias"`
	Address          string `json:"address"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`

	Extattr         `json:"extattr"`
	ExternalProfile `json:"external_profile"`
}

type ExternalProfile struct {
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
}

type Extattr struct {
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
}

type ReqListMembers struct {
	DepartmentID string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}

// IntoURLValues impl url.valuer
func (x ReqListMembers) IntoURLValues() url.Values {
	return url.Values{
		"department_id": {x.DepartmentID},
		// TODO FIX .........
		"fetch_child": {"false"},
	}
}

type RespListMembers struct {
	RespCommon
	UserList []struct {
		UserID     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
	} `json:"userlist"`
}

type ReqCreateDepartment struct {
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
	ID       int    `json:"id"`
}

type RespCreateDepartment struct {
	RespCommon
	ID int `json:"id"`
}

func (x ReqCreateDepartment) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

type ReqUpdateDepartment struct {
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
	ID       int    `json:"id"`
}

func (x ReqUpdateDepartment) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (x Member) IntoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
}

type ReqDepartmentDelete struct {
	DepartmentID string `json:"id"`
}

// IntoURLValues impl url.valuer
func (x ReqDepartmentDelete) IntoURLValues() url.Values {
	return url.Values{
		"id": {x.DepartmentID},
	}
}

type ReqListDepartments struct {
	DepartmentID string `json:"id"`
}

// IntoURLValues impl url.valuer
func (x ReqListDepartments) IntoURLValues() url.Values {
	return url.Values{
		"id": {x.DepartmentID},
	}
}

type RespListDepartments struct {
	RespCommon

	Department []Department `json:"department"`
}
