package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2SignUpdateParam{Sign: "sign"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sign.V2SignUpdate(p)
	fmt.Println(r.V2SignUpdateResult)
}
