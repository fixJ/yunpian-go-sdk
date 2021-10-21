package user

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type User struct {
	Apikey string
	Domain string
}

func (u User) V2UserGet(param param.V2UserGetParam) (result.Result, error) {
	apiUrl := u.Domain + "/v2/user/get.json"
	return request.HttpPostForm(apiUrl, param, "V2UserGetResult", u.Apikey)
}

func (u User) V2UserSet(param param.V2UserSetParam) (result.Result, error) {
	apiUrl := u.Domain + "/v2/user/set.json"
	return request.HttpPostForm(apiUrl, param, "V2UserSetResult", u.Apikey)
}