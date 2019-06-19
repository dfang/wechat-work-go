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

## Features (Refactoring ......)

<details>
<summary>通讯录管理 API</summary>

* [ ] [成员管理](https://work.weixin.qq.com/api/doc#90000/90135/90194)
    - [✓] 创建成员 (只能使用通讯录secret创建)
    - [✓] 读取成员 
    - [✓] 更新成员
    - [✓] 删除成员
    - [ ] 批量删除成员
    - [✓] 获取部门成员
    - [✓] 获取部门成员详情
    - [ ] userid与openid互换
    - [ ] 二次验证
    - [ ] 邀请成员

* [✓] [部门管理](https://work.weixin.qq.com/api/doc#90000/90135/90204)
    - [✓] 创建部门
    - [✓] 更新部门
    - [✓] 删除部门
    - [✓] 获取部门列表

* [ ] [标签管理](https://work.weixin.qq.com/api/doc#90000/90135/90209)

* [ ] [异步批量接口](https://work.weixin.qq.com/api/doc#90000/90135/90978)
    
</details>

<details>
<summary>身份认证</summary>

* [ ] [网页授权登录时获取访问用户身份](https://work.weixin.qq.com/api/doc#90000/90135/91023)
* [ ] [扫码授权登录时获取访问用户身份](https://work.weixin.qq.com/api/doc#90000/90135/91437)

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

* [✓] [发送应用消息](https://work.weixin.qq.com/api/doc#90000/90135/90236)
* [✓] 发送消息到群聊会话
    - [✓] 创建群聊会话
    - [✓] 修改群聊会话
    - [✓] 获取群聊会话
    - [✓] [应用推送消息](https://work.weixin.qq.com/api/doc#90000/90135/90248)

</details>


<details>
<summary>TODO(以下接口个人暂时用不到)</summary>

* [ ] 外部联系人管理
* [ ] OA 数据接口
* [ ] 企业支付
* [ ] 电子发票

</details>


## 对照API文档, 运行测试, 快速了解API

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

## Credits

直接从[Commits](https://github.com/xen0n/go-workwx/tree/5dbb164de258486669bbd9637d19e07124444d60)开始fork，获取access_token直接照搬的, 做了一些重构，添加了一些测试。感谢[xen0n](https://github.com/xen0n)。

目前api极不稳定, 快速迭代中 .....
