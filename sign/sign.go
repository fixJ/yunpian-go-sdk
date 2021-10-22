package sign

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type Sign struct {
	Apikey string
	Domain string
}

// V2SignAdd 添加签名
func (s Sign) V2SignAdd(param param.V2SignAddParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sign/add.json"
	return request.HttpPostForm(apiUrl, param, "V2SignAddResult", s.Apikey)
}

// V2SignGet 获取签名
func (s Sign) V2SignGet(param param.V2SignGetParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sign/get.json"
	return request.HttpPostForm(apiUrl, param, "V2SignGetResult", s.Apikey)
}

// V2SignUpdate 更新签名
func (s Sign) V2SignUpdate(param param.V2SignUpdateParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sign/update.json"
	return request.HttpPostForm(apiUrl, param, "V2SignUpdateResult", s.Apikey)
}