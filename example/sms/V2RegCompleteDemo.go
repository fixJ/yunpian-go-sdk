package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2RegCompleteParam{Mobile: "your mobile", Time: "time string"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2RegComplete(p)
	fmt.Println(r.V2RegCompleteResult)
}