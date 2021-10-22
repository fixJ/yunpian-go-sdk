package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsTplGetParam{TplId: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsTplGet(p)
	fmt.Println(r.V2VsmsTplGetResult)
}
