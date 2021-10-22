package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2SignAddParam{Sign: "sign"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Sign.V2SignAdd(p)
	fmt.Println(r.V2SignAddResult)
}
