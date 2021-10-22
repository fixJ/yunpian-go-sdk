package tpl

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type Tpl struct {
	Apikey	string
	Domain  string
}

// V2TplAdd 添加模版
func (t Tpl) V2TplAdd(param param.V2TplAddParam) (result.Result, error) {
	apiUrl := t.Domain + "/v2/tpl/add.json"
	return request.HttpPostForm(apiUrl, param, "V2TplAddResult", t.Apikey)
}

// V2TplGet 获取模版
func (t Tpl) V2TplGet(param param.V2TplGetParam) (result.Result, error) {
	apiUrl := t.Domain + "/v2/tpl/get.json"
	return request.HttpPostForm(apiUrl, param, "V2TplGetResult", t.Apikey)
}

// V2TplUpdate 更新模版
func (t Tpl) V2TplUpdate(param param.V2TplUpdateParam) (result.Result, error) {
	apiUrl := t.Domain + "/v2/tpl/update.json"
	return request.HttpPostForm(apiUrl, param, "V2TplUpdateResult", t.Apikey)
}

// V2TplDelete 删除模版
func (t Tpl) V2TplDelete(param param.V2TplDeleteParam) (result.Result, error) {
	apiUrl := t.Domain + "/v2/tpl/del.json"
	return request.HttpPostForm(apiUrl, param, "V2TplDeleteResult", t.Apikey)
}