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

// V2VoiceSend 发送语音验证码短信
func (v Voice) V2VoiceSend(param param.V2VoiceSendParam) (result.Result, error) {
	apiUrl := v.Domain + "/v2/voice/send.json"
	return request.HttpPostForm(apiUrl, param, "V2VoiceSendResult", v.Apikey)
}

// V2VoicePullStatus 获取语音验证码状态报告
func (v Voice) V2VoicePullStatus(param param.V2VoicePullStatusParam) (result.Result, error) {
	apiUrl := v.Domain + "/v2/voice/pull_status.json"
	return request.HttpPostForm(apiUrl, param, "V2VoicePullStatusResult", v.Apikey)
}

