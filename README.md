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
