package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2UserGetParam{}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.User.V2UserGet(p)
	fmt.Println(r.V2UserGetResult)
}
