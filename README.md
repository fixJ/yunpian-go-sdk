# yunpian-go-sdk

使用Golang封装的云片短信平台SDK

官方文档：https://www.yunpian.com/dev-doc

# 接口支持


| 接口名称               | 接口地址                             |
| ------------------------ | -------------------------------------- |
| 单条发送               | /v2/sms/single_send.json             |
| 批量发送相同内容       | /v2/sms/batch_send.json              |
| 指定模板单发           | /v2/sms/tpl_single_send.json         |
| 指定模板群发           | /v2/sms/tpl_batch_send.json          |
| 获取状态报告           | /v2/sms/pull_status.json             |
| 获取回复短信           | /v2/sms/pull_reply.json              |
| 添加模板               | /v2/tpl/add.json                     |
| 获取模板               | /v2/tpl/get.json                     |
| 修改模板               | /v2/tpl/update.json                  |
| 删除模板               | /v2/tpl/del.json                     |
| 添加签名               | /v2/sign/add.json                    |
| 获取签名               | /v2/sign/get.json                    |
| 修改签名               | /v2/sign/update.json                 |
| 查短信发送记录         | /v2/sms/get_record.json              |
| 注册成功回调           | /v2/sms/reg_complete.json            |
| 日账单导出             | /v2/sms/get_total_fee.json           |
| 语音验证码发送         | /v2/voice/send.json                  |
| 语音验证码状态报告获取 | /v2/voice/pull_status.json           |
| 批量发送超级短信       | /v2/vsms/tpl_batch_send.json         |
| 添加视频短信签名       | /v2/vsms/add_sign.json               |
| 查询视频短信签名       | /v2/vsms/search_sign.json            |
| 删除视频短信签名       | /v2/vsms/delete_sign.json            |
| 添加超级短信模板       | /v2/vsms/add_tpl.json                |
| 查询超级短信模板       | /v2/vsms/get_tpl.json                |
| 删除视频短信模板       | /v2/vsms/delete_tpl.json             |
| 查账户信息             | /v2/user/get.json                    |
| 修改账号信息           | /v2/user/set.json                    |
| 创建短链接             | /v2/short_url/shorten.json           |
| 短链访问统计查询       | /v2/short_url/time_series_stats.json |

# 获取SDK

`go get github.com/fixJ/yunpian-go-sdk`

# 快速接入

下面这个例子可以实现短信单发功能

```go
package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
    p := param.V2SingleSendParam{Mobile: "", Text: ""}
    c := client.InitDefaultClient("your apikey")
    r, _ := c.Sms.V2SingleSend(p)
    fmt.Println(r)
}
```

# SDK接入流程

1. 获取SDK
2. 初始化client
3. 构造请求参数对象
4. 通过client调用对应的方法
5. 获取返回值

## 获取SDK

使用go get github.com/fixJ/yunpian-go-sdk下载SDK

## 初始化client

调用client package下的InitDefaultClient/InitCustomClient来初始化client

需要传入apikey；InitCustomClient还需额外传入domain，这个参数会作为内部接口请求的域名

example:

```go
// 默认的初始化方法
c1 := client.InitDefaultClient("your apikey")

// 手动传入域名方法
c2 := client.InitCustomClient("your apikey", "https://sms.yunpian.com")

```

### domain说明

默认的InitDefaultClient初始化的client使用的域名是https://sms.yunpian.com

暴露出来给用户设置的好处：

1. 用户环境无法支持https
2. 用户可以根据网络环境使用云片国内的域名sms.yunpian.com或者国外的域名us.yunpian.com

## 构造请求参数对象

每个接口的请求参数类型都放在param package下

根据你要使用的接口，创建对应的param

example:

```go
// 创建了一个用于V2单发接口的param

p := param.V2SingleSendParam{Mobile: "", Text: ""}
```

## 通过client调用对应的方法

所有接口方法已经分类放在对应的package下，这些方法都挂载一个对应的type struct上


| package  | 功能         |
| ---------- | -------------- |
| sms      | 短信发送功能 |
| sign     | 签名功能     |
| tpl      | 模版功能     |
| user     | 用户功能     |
| voice    | 语音功能     |
| vsms     | 视频短信功能 |
| shortUrl | 短链接功能   |

YunPianClient结构体中挂载了所有的接口方法的type struct，所以通过client可以访问到接口对应的方法

example:

```go
// 调用client下Sms struct中的V2SingleSend方法，即短信单发接口
// 传入的是V2SingleSend对应的参数

c := client.InitDefaultClient("your apikey")
p := param.V2SingleSendParam{Mobile: "", Text: ""}
r, err := c.Sms.V2SingleSend(p)
```

## 获取返回值

每个接口方法的返回值类型放在result package中

Result结构体中Status字段表示了此次请求接口是否成功，成功为0，失败为-1，判断标准是接口返回的http状态码是否为200

接口请求失败后的返回值统一封装在Result的FailResult中

其余的字段是所有接口正常返回的结构体类型，虽然Result中封装了所有接口正常返回的结构体类型，但只有调用的接口方法所对应的类型会被赋值

调用过程中所有的内部错误都由err返回

example:

```go
//调用V2SingleSend方法，并通过Result中的V2SingleSendResult获取其返回值

c := client.InitDefaultClient("your apikey")
p := param.V2SingleSendParam{Mobile: "", Text: ""}
r, err := c.Sms.V2SingleSend(p)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(r.V2SingleSendResult)
}

```

## Example

所有接口均提供示例，请参考 https://github.com/fixJ/yunpian-go-sdk/tree/master/example
