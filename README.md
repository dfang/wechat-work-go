# wechat-work-go

[![Go Report Card](https://goreportcard.com/badge/github.com/dfang/wechat-work-go)](https://goreportcard.com/report/github.com/dfang/wechat-work-go)
[![GoDoc](http://godoc.org/github.com/dfang/wechat-work-go?status.svg)](http://godoc.org/github.com/dfang/wechat-work-go)

```go
import (
    WechatWork "github.com/dfang/wechat-work-go" // package WechatWork
)
```

## Features

* [x] access token 刷新
* [ ] 通讯录管理
* [ ] 外部联系人管理
* [ ] 应用管理
* [x] 消息发送 (**部分支持**，见下)
* [ ] 素材管理

<details>
<summary>通讯录管理 API</summary>

* [ ] 成员管理
    - [ ] 创建成员
    - [x] 读取成员 *NOTE: 成员对外信息暂未实现*
    - [ ] 更新成员
    - [ ] 删除成员
    - [ ] 批量删除成员
    - [ ] 获取部门成员
    - [ ] 获取部门成员详情
    - [ ] userid与openid互换
    - [ ] 二次验证
    - [ ] 邀请成员
* [ ] 部门管理
    - [ ] 创建部门
    - [ ] 更新部门
    - [ ] 删除部门
    - [ ] 获取部门列表
* [ ] 标签管理
    - [ ] 创建标签
    - [ ] 更新标签名字
    - [ ] 删除标签
    - [ ] 获取标签成员
    - [ ] 增加标签成员
    - [ ] 删除标签成员
    - [ ] 获取标签列表
* [ ] 异步批量接口
    - [ ] 增量更新成员
    - [ ] 全量覆盖成员
    - [ ] 全量覆盖部门
    - [ ] 获取异步任务结果
* [ ] 通讯录回调通知
    - [ ] 成员变更通知
    - [ ] 部门变更通知
    - [ ] 标签变更通知
    - [ ] 异步任务完成通知

</details>

<details>
<summary>外部联系人管理 API</summary>

* [ ] 离职成员的外部联系人再分配
* [ ] 成员对外信息
* [ ] 获取外部联系人详情

</details>

<details>
<summary>应用管理 API</summary>

* [ ] 获取应用
* [ ] 设置应用
* [ ] 自定义菜单
    - [ ] 创建菜单
    - [ ] 获取菜单
    - [ ] 删除菜单

</details>

<details>
<summary>消息发送 API</summary>

* [x] 发送应用消息
* [ ] 接收消息
* [x] 发送消息到群聊会话
    - [ ] 创建群聊会话
    - [ ] 修改群聊会话
    - [ ] 获取群聊会话
    - [x] 应用推送消息

### 消息类型

* [x] 文本消息
* [ ] 图片消息
* [ ] 语音消息
* [ ] 视频消息
* [ ] 文件消息
* [ ] 文本卡片消息
* [ ] 图文消息
* [ ] 图文消息（mpnews）
* [x] markdown消息

</details>

## 运行测试, 快速上手

```
export CORP_ID=xxxxxx
export CORP_SECRET=yyyyyy
export AGENT_ID=zzzzzz

go test -v access_token_test.go
```

## Credits

直接从[Commits](https://github.com/xen0n/go-workwx/tree/5dbb164de258486669bbd9637d19e07124444d60)开始fork，获取access_token直接照搬的, 做了一些重构，添加了一些测试。感谢[xen0n](https://github.com/xen0n)。

目前api极不稳定, 快速迭代中 .....

## Notes
### 目前无法通过api创建应用和删除群聊会话, 因为tx没有提供这样的api

参考 `go test -v agent_test.go chat_test.go`

### 如果快速找到corpid, corp_secret, agent_id ?

corp_id: 管理后台-> 我的企业 -> 企业ID（在最底部)  
corp_secret和agent_id去自己创建的应用详情页面找，(管理后台-> 应用与小程序 -> 点击应用进详情，能看到AgentId 和 secret)


以下信息从企业客服聊天得知:  
    
>    secret就是等于corpsecret，目前企业微信整体来说有3个地方是有secret，分别是应用、外部联系人、管理工具-通讯录同步助手
    参考文档：https://work.weixin.qq.com/api/doc#10013

>    只有api创建的才有chatid, 其他途径是没有办法去获取的 https://work.weixin.qq.com/api/doc#13288

总结:   
    
>    所以corp_secret 其实不应该叫corp_secret, 叫app_secret 或 agent_secret 比较合适。

## License

* [MIT](./LICENSE)
