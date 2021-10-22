package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsSignDeleteParam{
		Sign: "",
	}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsSignDelete(p)
	fmt.Println(r.V2VsmsSignDeleteResult)
}
