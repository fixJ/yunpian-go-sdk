package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2ShortUrlStatsParam{Sid: "", StartTime: "2021-01-01 00:00:00", EndTime: "2021-02-01 00:00:00"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.ShortUrl.V2ShortUrlStats(p)
	fmt.Println(r.V2ShortUrlStatsResult)
}
