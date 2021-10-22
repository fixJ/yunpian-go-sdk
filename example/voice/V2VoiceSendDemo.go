package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VoiceSendParam{Mobile: "", Code: "1234"}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Voice.V2VoiceSend(p)
	fmt.Println(r.V2VoiceSendResult)
}
