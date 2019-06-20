package contact

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

type RespListMembers struct {
	RespCommon
	UserList []struct {
		UserID     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
	} `json:"userlist"`
}

type RespListMembers2 struct {
	RespCommon

	UserList []struct {
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
		Status         int    `json:"status"`
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

type ReqUpdateDepartment struct {
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
	ID       int    `json:"id"`
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
}

type RespListDepartments struct {
	RespCommon

	Department []Department `json:"department"`
}

type ReqBatchDeleteMembers struct {
	UserIDList []string `json:"useridlist"`
}

// RespCommon Comman Response Struct
type RespCommon struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
