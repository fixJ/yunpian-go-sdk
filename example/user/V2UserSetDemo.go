package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2UserSetParam{AlarmBalance: "200", EmergencyContact: "", EmergencyMobile: ""}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.User.V2UserSet(p)
	fmt.Println(r.V2UserSetResult)
}
