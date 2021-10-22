package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2PullStatusParam{PageSize: "10"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2PullStatus(p)
	fmt.Println(r.V2PullStatusResult)
}