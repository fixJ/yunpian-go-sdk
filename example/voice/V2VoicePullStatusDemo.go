package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VoicePullStatusParam{PageSize: "10"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Voice.V2VoicePullStatus(p)
	fmt.Println(r.V2VoicePullStatusResult)
}
