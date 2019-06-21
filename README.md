# wechat-work-go

[![Go Report Card](https://goreportcard.com/badge/github.com/dfang/wechat-work-go)](https://goreportcard.com/report/github.com/dfang/wechat-work-go)
[![GoDoc](http://godoc.org/github.com/dfang/wechat-work-go?status.svg)](http://godoc.org/github.com/dfang/wechat-work-go)
[![Travis](https://travis-ci.com/dfang/wechat-work-go.svg?branch=refactor)](https://travis-ci.com/dfang/wechat-work-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/a054e30c788eb3a693ac/maintainability)](https://codeclimate.com/github/dfang/wechat-work-go/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a054e30c788eb3a693ac/test_coverage)](https://codeclimate.com/github/dfang/wechat-work-go/test_coverage)
[![Build status](https://badge.buildkite.com/988dd09408df5bf2bb39403f4f3c10f50e05b84846ba6f6791.svg)](https://buildkite.com/curry/wechat-work-go)

```go
import (
    wechatwork "github.com/dfang/wechat-work-go" // package wechatwork
)
```

## Features

<details>
<summary>通讯录管理 API</summary>

* [✓] [成员管理](https://work.weixin.qq.com/api/doc#90000/90135/90194)
    - [✓] 创建成员 (只能使用通讯录secret创建)
    - [✓] 读取成员 
    - [✓] 更新成员
    - [✓] 删除成员
    - [✓] 批量删除成员
    - [✓] 获取部门成员
    - [✓] 获取部门成员详情
    - [✓] userid与openid互换
    - [✓] 二次验证
    - [✓] 邀请成员

* [✓] [部门管理](https://work.weixin.qq.com/api/doc#90000/90135/90204)
    - [✓] 创建部门
    - [✓] 更新部门
    - [✓] 删除部门
    - [✓] 获取部门列表

* [ ] [标签管理](https://work.weixin.qq.com/api/doc#90000/90135/90209)
    - [ ] 创建标签
    - [ ] 更新标签名字
    - [ ] 删除标签
    - [ ] 获取标签成员
    - [ ] 增加标签成员
    - [ ] 删除标签成员
    - [ ] 获取标签列表

* [ ] [异步批量接口](https://work.weixin.qq.com/api/doc#90000/90135/90978)
    - [ ] 增量更新成员
    - [ ] 全量覆盖成员
    - [ ] 全量覆盖部门
    - [ ] 获取异步任务结果

* [ ] [通讯录回调通知](https://work.weixin.qq.com/api/doc#90000/90135/90966)


</details>

<details>
<summary>身份认证</summary>

* [✓] [网页授权登录时获取访问用户身份](https://work.weixin.qq.com/api/doc#90000/90135/91023)
* [✓] [扫码授权登录时获取访问用户身份](https://work.weixin.qq.com/api/doc#90000/90135/91437)

</details>

<details>
<summary>应用管理</summary>

* [✓] 获取应用
* [✓] 设置应用
* [✓] [自定义菜单](https://work.weixin.qq.com/api/doc#90000/90135/90230)
    - [✓] 创建菜单
    - [✓] 获取菜单
    - [✓] 删除菜单

</details>

<details>
<summary>消息推送</summary>

* [✓] [推送消息到应用](https://work.weixin.qq.com/api/doc#90000/90135/90236)
* [✓] 发送消息到群聊会话
    - [✓] 创建群聊会话
    - [✓] 修改群聊会话
    - [✓] 获取群聊会话
    - [✓] [推送消息到群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90248)
* [ ] 互联企业消息推送
    - [ ] 发送应用消息

</details>

<details>
<summary>素材管理</summary>

* [ ] 素材管理
    - [ ] 上传临时素材
    - [ ] 上传永久图片
    - [ ] 获取临时素材
    - [ ] 获取高清语音素材

</details>

<details>
<summary>TODO(以下接口个人暂时用不到， 暂时不开发)</summary>

* [ ] 外部联系人管理
* [ ] OA 数据接口
* [ ] 企业支付
* [ ] 电子发票

</details>


## 使用

企业微信是分应用的， 一个企业（一个corpid）中有多个app （对应的一套corpsecret，agentid）  

所以使用的时候先创建一个企业的client，然后用这个client创建不同的app对象  

```go
import (
    wechatwork "github.com/dfang/wechat-work-go" // package wechatwork
)

corpID := os.Getenv("CORP_ID")
corpSecret := os.Getenv("CORP_SECRET")
agentID, _ := strconv.ParseInt(os.Getenv("AGENT_ID"), 10, 64)

client := wechatwork.New(corpID)
app = client.NewApp(corpSecret, agentID)
    
```

要使用哪个模块的功能，就创建哪个模块的实例


```
import "github.com/dfang/wechat-work-go/message"

msg := message.WithApp(app)

msg.SendAppChatMessage(....)
```





## 开发 

对照API文档, 运行测试, 快速了解API

```
export CORP_ID=xxxxxx
export CORP_SECRET=yyyyyy
export AGENT_ID=zzzzzz

运行单个测试，比如只运行access_token_test.go的测试
ginkgo -v -focus='access_token'

或者
go test -v wechat_work_go_suite_test.go client_test.go

watch mode  
ginkgo watch -v -focus='access_token*'
```
