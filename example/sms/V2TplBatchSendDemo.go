package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2TplBatchSendParam{Mobile: "your mobile", TplId: "your tpl_id", TplValue: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2TplBatchSend(p)
	fmt.Println(r.V2TplBatchSendResult)
}