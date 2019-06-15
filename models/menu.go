package models

// 创建自定义菜单
type ReqMenuCreate struct {
	Button []struct {
		Name      string `json:"name"`
		Type      string `json:"type,omitempty"`
		Key       string `json:"key,omitempty"`
		SubButton []struct {
			Type      string        `json:"type"`
			Name      string        `json:"name"`
			Key       string        `json:"key"`
			URL       string        `json:"url,omitempty"`
			SubButton []interface{} `json:"sub_button"`
		} `json:"sub_button,omitempty"`
	} `json:"button"`
}
