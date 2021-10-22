package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2TplAddParam{TplContent: "tpl content"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Tpl.V2TplAdd(p)
	fmt.Println(r.V2TplAddResult)
}
