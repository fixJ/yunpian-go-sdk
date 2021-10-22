package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsSignSearchParam{
		Sign: "",
		PageSize: "10",
		PageNum: "1",
	}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsSignSearch(p)
	fmt.Println(r.V2VsmsSignSearchResult)
}
