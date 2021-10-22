package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2GetTotalFeeParam{Date: "date of total fee"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sms.V2GetTotalFee(p)
	fmt.Println(r.V2GetTotalFeeResult)
}