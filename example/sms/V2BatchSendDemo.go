package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2BatchSendParam{Mobile: "your mobile", Text: "your text"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2BatchSend(p)
	fmt.Println(r.V2BatchSendResult)
}