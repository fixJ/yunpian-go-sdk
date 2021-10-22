package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2TplDeleteParam{TplId: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Tpl.V2TplDelete(p)
	fmt.Println(r.V2TplDeleteResult)
}
