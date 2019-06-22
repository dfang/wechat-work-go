// Package wechatwork 企业微信api的封装
//
// https://work.weixin.qq.com/api/doc#90000/90003/90556
//
// 商户版见 https://github.com/dfang/wechat-work-go-sp
package wechatwork

import (
	"os"
	"sync"

	"github.com/dfang/wechat-work-go/cache"
	"gopkg.in/resty.v1"
)

// WechatWork 企业微信客户端
type WechatWork struct {
	// CorpID 企业 ID，必填
	CorpID string
}

// App 企业微信客户端（分应用）
type App struct {
	*WechatWork

	// CorpSecret 应用的凭证密钥，其实应该叫AgentSecret更好，必填
	CorpSecret string

	// AgentID 应用 ID，必填
	AgentID int64

	AccessToken string

	// Token     string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`

	accessTokenLock *sync.RWMutex

	Cache cache.Cache
}

// New 构造一个 WechatWork 对象，需要提供企业 ID
//
// 通常，要使用wechat-work-go, 你需要先创建一个 WechatWork 的对象，
// 接着以此对象调用 WithApp 创建一个app
// 然后就可以以app对象调用API了
//
// 企业微信API是分应用的
//
// 简单来说就是 a 应用的CORP_SECRET 和 AGENT_ID 获取的access_token 是不能操作 b应用的
//
// 	CORP_ID 去企业微信管理中的 我的企业 最底部
// 	CORP_SECRET 其实是应用的secret，个人感觉应该叫AgentSecret 或 AppSecret 更合适, 但因为api接口和官方文档叫corpSecret，所以不改变
// 	AGENT_ID 应用的ID， CORP_SECRET 和 AGENT_ID 都去应用的详情页面找
//
// 示例代码:
//
// 		corpID := os.Getenv("CORP_ID")
// 		corpSecret := os.Getenv("CORP_SECRET")
// 		agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)
// 		corp := wechatwork.New(corpID)
//		app = corp.NewApp(corpSecret, agentID)
//
func New(corpID string) *WechatWork {
	return &WechatWork{
		CorpID: corpID,
	}
}

// NewApp 构造本企业下某自建 app 的对象
//
// 企业微信暂未提供创建 app 的 api, 创建应用需要去企业微信的管理后台中
//
func (corp *WechatWork) NewApp(corpSecret string, agentID int64) *App {
	return &App{
		WechatWork: corp,

		CorpSecret:      corpSecret,
		AgentID:         agentID,
		AccessToken:     "",
		accessTokenLock: &sync.RWMutex{},
		Cache:           cache.NewMemory(),
	}
}

// newDefaultRestyClient 返回一个resty 的client
func newDefaultRestyClient() *resty.Client {
	client := resty.New()
	client.SetDebug(os.Getenv("DEBUG") == "true")
	client.SetLogger(os.Stdout)
	client.SetHostURL("https://qyapi.weixin.qq.com")
	return client
}

// SimpleGet 一切get请求的api调用可使用此方法
//
// 企业微信中，获取操作和删除都是GET请求
//
// just like resty.SetReult(&respObj).Get(url)
//
// note: url must be full, if you're consuming this library, you don't need this method, have to export this method for other packages to call
func (app *App) SimpleGet(url string, respObj interface{}) error {
	resp, err := newDefaultRestyClient().R().
		SetHeader("Accept", "application/json").
		SetResult(&respObj).
		Get(url)

	if err != nil {
		panic(err)
	}
	return nil
}

// SimplePost 一切Post请求的api调用使用此方法
//
// 企业微信中，删除操作一般都是GET请求，更新操作、批量删除成员是POST请求，没有PUT、PATCH、DELETE
//
// just like resty.SetBody(b).SetReult(&respObj).Post(url)
//
// resty can Automatically marshal and unmarshal
//
// note: url must be full, if you're consuming this library, you don't need this method, have to export this method for other packages to call
func (app *App) SimplePost(url string, body interface{}, respObj interface{}) error {
	resp, err := newDefaultRestyClient().R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&respObj).
		Post(url)

	if err != nil {
		panic(err)
	}
	return nil
}

// RespCommon Comman Response Struct
type RespCommon struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
//
// 实现依据: https://work.weixin.qq.com/api/doc#10013
//
// > 企业微信所有接口，返回包里都有errcode、errmsg。
//
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
//
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *RespCommon) IsOK() bool {
	return x.ErrCode == 0
}
