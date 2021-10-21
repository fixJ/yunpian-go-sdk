package voice

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type Voice struct {
	Apikey string
	Domain string
}

func (v Voice) V2VoiceSend(param param.V2VoiceSendParam) (result.Result, error) {
	apiUrl := v.Domain + "/v2/voice/send.json"
	return request.HttpPostForm(apiUrl, param, "V2VoiceSendResult", v.Apikey)
}

func (v Voice) V2VoicePullStatus(param param.V2VoicePullStatusParam) (result.Result, error) {
	apiUrl := v.Domain + "/v2/voice/pull_status.json"
	return request.HttpPostForm(apiUrl, param, "V2VoicePullStatusResult", v.Apikey)
}

