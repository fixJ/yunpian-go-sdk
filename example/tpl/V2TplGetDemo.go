package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2TplGetParam{TplId: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Tpl.V2TplGet(p)
	fmt.Println(r.V2TplGetResult)
}
