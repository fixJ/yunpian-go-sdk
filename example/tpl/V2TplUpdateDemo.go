package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2TplUpdateParam{TplId: "", TplContent: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Tpl.V2TplUpdate(p)
	fmt.Println(r.V2TplUpdateResult)
}
