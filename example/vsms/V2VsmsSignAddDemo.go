package main

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	p := param.V2VsmsSignAddParam{
		Sign: "",
		CertificateFileContents: "",
		CertificateFileSuffix: "",
		LicenseFileContents: "",
		LicenseFileSuffix: "",
	}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsSignAdd(p)
	fmt.Println(r.V2VsmsSignAddResult)
}
