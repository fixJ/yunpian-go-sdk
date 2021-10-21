package sms

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type Sms struct {
	Apikey string
	Domain string
}

func (s Sms) V2SingleSend(param param.V2SingleSendParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/single_send.json"
	return request.HttpPostForm(apiUrl, param, "V2SingleSendResult", s.Apikey)
}

func (s Sms) V2BatchSend(param param.V2BatchSendParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/batch_send.json"
	return request.HttpPostForm(apiUrl, param, "V2BatchSendResult", s.Apikey)
}

func (s Sms) V2TplSingleSend(param param.V2TplSingleSendParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/tpl_single_send.json"
	return request.HttpPostForm(apiUrl, param, "V2TplSingleSendResult", s.Apikey)
}

func (s Sms) V2TplBatchSend(param param.V2TplBatchSendParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/tpl_batch_send.json"
	return request.HttpPostForm(apiUrl, param, "V2TplBatchSendResult", s.Apikey)
}

func (s Sms) V2PullStatus(param param.V2PullStatusParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/pull_status.json"
	return request.HttpPostForm(apiUrl, param, "V2PullStatusResult", s.Apikey)
}

func (s Sms) V2PullReply(param param.V2PullReplyParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/pull_reply.json"
	return request.HttpPostForm(apiUrl, param, "V2PullReplyResult", s.Apikey)
}

func (s Sms) V2GetRecord(param param.V2GetRecordParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/get_record.json"
	return request.HttpPostForm(apiUrl, param, "V2GetRecordResult", s.Apikey)
}

func (s Sms) V2RegComplete(param param.V2RegCompleteParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/reg_complete.json"
	return request.HttpPostForm(apiUrl, param, "V2RegCompleteResult", s.Apikey)
}

func (s Sms) V2GetTotalFee(param param.V2GetTotalFeeParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/sms/get_total_fee.json"
	return request.HttpPostForm(apiUrl, param, "V2GetTotalFeeResult", s.Apikey)
}