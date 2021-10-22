package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsTplDeleteParam{TplId: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsTplDelete(p)
	fmt.Println(r.V2VsmsTplDeleteResult)
}
