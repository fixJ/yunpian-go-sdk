package client

import (
	"github.com/fixJ/yunpian-go-sdk/constant"
	"github.com/fixJ/yunpian-go-sdk/shortUrl"
	"github.com/fixJ/yunpian-go-sdk/sign"
	"github.com/fixJ/yunpian-go-sdk/sms"
	"github.com/fixJ/yunpian-go-sdk/tpl"
	"github.com/fixJ/yunpian-go-sdk/user"
	"github.com/fixJ/yunpian-go-sdk/voice"
)

type YunPianClient struct {
	apikey string
	domain string
	Sms    sms.Sms
	Tpl    tpl.Tpl
	Sign   sign.Sign
	Voice voice.Voice
	User  user.User
	ShortUrl shortUrl.ShortUrl
}


func InitDefaultClient(apikey string) YunPianClient {
	dd := constant.DefaultDomain
	return YunPianClient{
		apikey: apikey,
		domain: dd,
		Sms: sms.Sms{Apikey: apikey, Domain: dd},
		Tpl: tpl.Tpl{Apikey: apikey, Domain: dd},
		Sign: sign.Sign{Apikey: apikey, Domain: dd},
		Voice: voice.Voice{Apikey: apikey, Domain: dd},
		User: user.User{Apikey: apikey, Domain: dd},
		ShortUrl: shortUrl.ShortUrl{Apikey: apikey, Domain: dd},
	}
}

func InitCustomClient(apikey string, domain string) YunPianClient {
	return YunPianClient{
		apikey: apikey,
		domain: domain,
		Sms: sms.Sms{Apikey: apikey, Domain: domain},
		Tpl: tpl.Tpl{Apikey: apikey, Domain: domain},
		Sign: sign.Sign{Apikey: apikey, Domain: domain},
		Voice: voice.Voice{Apikey: apikey, Domain: domain},
		User: user.User{Apikey: apikey, Domain: domain},
		ShortUrl: shortUrl.ShortUrl{Apikey: apikey, Domain: domain},
	}
}
