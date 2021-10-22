package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2SingleSendParam{Mobile: "your mobile", Text: "your text"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2SingleSend(p)
	fmt.Println(r.V2SingleSendResult)
}