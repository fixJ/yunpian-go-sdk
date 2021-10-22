package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2GetRecordParam{PageSize: "10"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2GetRecord(p)
	fmt.Println(r.V2GetRecordResult)
}