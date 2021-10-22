package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsTplBatchSendParam{TplId: "", Mobile: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsTplBatchSend(p)
	fmt.Println(r.V2VsmsTplBatchSendResult)
}
