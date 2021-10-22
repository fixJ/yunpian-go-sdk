package vsms

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type Vsms struct {
	Apikey string
	Domain string
}

func (v Vsms) V2VsmsTplBatchSend(param param.V2VsmsTplBatchSendParam) (result.Result, error) {
	apiUrl := v.Domain + "/v2/vsms/tpl_batch_send.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsTplBatchSendResult", v.Apikey)
}

func (v Vsms) V2VsmsSignAdd(param param.V2VsmsSignAddParam) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/add_sign.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsSignAddResult", v.Apikey)
}

func (v Vsms) V2VsmsSignSearch(param param.V2VsmsSignSearchParam) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/search_sign.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsSignSearchResult", v.Apikey)
}

func (v Vsms) V2VsmsSignDelete(param param.V2VsmsSignDeleteParam) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/delete_sign.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsSignDeleteResult", v.Apikey)
}

func (v Vsms) V2VsmsTplAdd(param param.V2VsmsTplAddParam, materialPath string) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/add_tpl.json"
	return request.HttpPostMultiPartForm(apiUrl, param, "V2VsmsTplAddResult", v.Apikey, materialPath)
}

func (v Vsms) V2VsmsTplGet(param param.V2VsmsTplGetParam) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/get_tpl.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsTplGetResult", v.Apikey)
}

func (v Vsms) V2VsmsTplDelete(param param.V2VsmsTplDeleteParam) (result.Result, error) {
	apiUrl := "https://vsms.yunpian.com" + "/v2/vsms/delete_tpl.json"
	return request.HttpPostForm(apiUrl, param, "V2VsmsTplDeleteResult", v.Apikey)
}