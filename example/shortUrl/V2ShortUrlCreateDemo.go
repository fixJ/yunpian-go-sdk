package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2ShortUrlCreateParam{LongUrl: "https://yunpian.com", Name: "test"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.ShortUrl.V2ShortUrlCreate(p)
	fmt.Println(r.V2ShortUrlCreateResult)
}
