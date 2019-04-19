package WechatWork

// CreateAppChat 创建群聊会话
func (c *App) CreateAppChat(req ReqAppChatCreate) (RespAppChatCreate, error) {
	// /cgi-bin/appchat/create?access_token=ACCESS_TOKEN
	apiPath := "/cgi-bin/appchat/create"
	var resp RespAppChatCreate
	err := c.executeQyapiJSONPost(apiPath, req, &resp, true)
	if err != nil {
		return RespAppChatCreate{}, err
	}

	return resp, nil
	// return errors.New("not implemented")
}

// UpdateAppChat 修改群聊会话
func (c *App) UpdateAppChat(req ReqAppChatUpdate) (RespCommon, error) {
	// /cgi-bin/appchat/update?access_token=ACCESS_TOKEN
	apiPath := "/cgi-bin/appchat/update"
	var resp RespCommon
	err := c.executeQyapiJSONPost(apiPath, req, &resp, true)
	if err != nil {
		return RespCommon{}, err
	}

	return resp, nil
	// return errors.New("not implemented")
}

// GetAppChat 获取群聊会话
func (c *App) GetAppChat(chatid string) (RespAppChatGet, error) {
	// /cgi-bin/appchat/get?access_token=ACCESS_TOKEN&chatid=CHATID
	apiPath := "/cgi-bin/appchat/get"

	req := ReqAppChatGet{
		AccessToken: c.AccessToken,
		ChatID:      chatid,
	}
	var resp RespAppChatGet
	err := c.executeQyapiGet(apiPath, req, &resp, true)

	if err != nil {
		return RespAppChatGet{}, err
	}

	return resp, nil
}

// SendAppChat 发送群聊会话(应用推送消息)
func (c *App) SendAppChat(appchat bodyer) (RespCommon, error) {
	// /cgi-bin/appchat/send?access_token=ACCESS_TOKEN
	apiPath := "/cgi-bin/appchat/send"
	// req := sendable
	var resp RespCommon
	err := c.executeQyapiJSONPost(apiPath, appchat, &resp, true)
	if err != nil {
		// TODO: error_chain
		return RespCommon{}, err
	}
	return resp, nil
}
