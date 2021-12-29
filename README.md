# yunpian-go-sdk

Yunpian SMS platform SDK by Golang

Official API Document：https://www.yunpian.com/dev-doc

[SDK 中文文档](https://github.com/fixJ/yunpian-go-sdk/blob/master/README.zh-CN.md)

# API Support

| 接口名称               | 接口地址                             |
| ------------------------ | -------------------------------------- |
| single send               | /v2/sms/single_send.json             |
| batch send       | /v2/sms/batch_send.json              |
| template single send           | /v2/sms/tpl_single_send.json         |
| template batch send           | /v2/sms/tpl_batch_send.json          |
| pull report           | /v2/sms/pull_status.json             |
| pull reply           | /v2/sms/pull_reply.json              |
| add template               | /v2/tpl/add.json                     |
| get template               | /v2/tpl/get.json                     |
| update template               | /v2/tpl/update.json                  |
| delete template               | /v2/tpl/del.json                     |
| add sign               | /v2/sign/add.json                    |
| get sign               | /v2/sign/get.json                    |
| update sign               | /v2/sign/update.json                 |
| get sms record         | /v2/sms/get_record.json              |
| register callback           | /v2/sms/reg_complete.json            |
| export total fee             | /v2/sms/get_total_fee.json           |
| voice sms send         | /v2/voice/send.json                  |
| pull voice sms report | /v2/voice/pull_status.json           |
| template batch send video sms       | /v2/vsms/tpl_batch_send.json         |
| add video sign       | /v2/vsms/add_sign.json               |
| get video sign       | /v2/vsms/search_sign.json            |
| delete video sign       | /v2/vsms/delete_sign.json            |
| add video template       | /v2/vsms/add_tpl.json                |
| get video template       | /v2/vsms/get_tpl.json                |
| delete video template       | /v2/vsms/delete_tpl.json             |
| get account info             | /v2/user/get.json                    |
| update account info           | /v2/user/set.json                    |
| create short link             | /v2/short_url/shorten.json           |
| get short link click info       | /v2/short_url/time_series_stats.json |

# Get SDK

`go get github.com/fixJ/yunpian-go-sdk`

# Quick Start

The following is an example of a single SMS message

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

# SDK Use steps

1. Get SDK
2. Init a client
3. Build request param object
4. Use client call the function
5. Get the response

## Get SDK

go get github.com/fixJ/yunpian-go-sdk download SDK

## Init client

call client package下的InitDefaultClient/InitCustomClient to init a client

need incoming an apikey；InitCustomClient need incoming a domain，the domain will be used as the api prefix

example:

```go
// default init function
c1 := client.InitDefaultClient("your apikey")

// incoming the domain by manual
c2 := client.InitCustomClient("your apikey", "https://sms.yunpian.com")

```

### Domain 

The default InitDefaultClient Initializes the client using the domain name https://sms.yunpian.com



The benefits of exposing Settings to users:



1. The user environment does not support HTTPS

2. Users can use the domestic domain name sms.yunpian.com or the foreign domain name us.yunpian.com according to the network environment

## Build request param object

The request parameter types for each api are placed under the Param Package



Create a param based on the api you want to use

example:

```go
// A param was created for the V2 single send api


p := param.V2SingleSendParam{Mobile: "", Text: ""}
```

## Use client call the function

All interface methods have been categorized under the corresponding package, and these methods are mounted on a corresponding Type struct

| package  | function         |
| ---------- | -------------- |
| sms      | send message |
| sign     | message's sign     |
| tpl      | message's template     |
| user     | platform user     |
| voice    | voice message     |
| vsms     | video message |
| shortUrl | short link   |

YunPianClient struct has mounted all the type struct of the interface method, so the corresponding method of the interface can be accessed through client
example:

```go
// Call V2SingleSend in the Sms struct of client

// V2SingleSend is passed in as a parameter

c := client.InitDefaultClient("your apikey")
p := param.V2SingleSendParam{Mobile: "", Text: ""}
r, err := c.Sms.V2SingleSend(p)
```

## Get response

The return value type for each interface method is placed in the Result Package



The Status field in the Result structure indicates whether the request is successful. The value of success is 0, and the value of failure is -1. The criterion is whether the HTTP Status code returned by the interface is 200



The returned value after an interface request fails is encapsulated in FailResult of Result



The remaining fields are the normal struct types returned by all interfaces. Although Result encapsulates the normal struct types returned by all interfaces, only the type of the interface method called is assigned



All internal errors during the call are returned by err

example:

```go
//Call the V2SingleSend method and get its return value from the V2SingleSendResult in Result


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

All API are provided with examples for your reference https://github.com/fixJ/yunpian-go-sdk/tree/master/example
