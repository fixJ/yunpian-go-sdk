package result

import (
	"encoding/json"
)

func Format(data string, r *Result, resType string) error {
	var err error = nil
	if resType == "FailResult" {
		err = json.Unmarshal([]byte(data), &((*r).FailResult))
	} else if resType == "V2SingleSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2SingleSendResult))
	} else if resType == "V2BatchSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2BatchSendResult))
	} else if resType == "V2TplSingleSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplSingleSendResult))
	} else if resType == "V2TplBatchSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplBatchSendResult))
	} else if resType == "V2PullStatusResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2PullStatusResult))
	} else if resType == "V2PullReplyResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2PullReplyResult))
	} else if resType == "V2TplAddResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplAddResult))
	} else if resType == "V2TplGetResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplGetResult))
	} else if resType == "V2TplUpdateResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplUpdateResult))
	} else if resType == "V2TplDeleteResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2TplDeleteResult))
	} else if resType == "V2SignAddResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2SignAddResult))
	} else if resType == "V2SignGetResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2SignGetResult))
	} else if resType == "V2SignUpdateResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2SignUpdateResult))
	} else if resType == "V2GetRecordResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2GetRecordResult))
	} else if resType == "V2RegCompleteResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2RegCompleteResult))
	} else if resType == "V2GetTotalFeeResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2GetTotalFeeResult))
	} else if resType == "V2VoiceSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VoiceSendResult))
	} else if resType == "V2VoicePullStatusResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VoicePullStatusResult))
	} else if resType == "V2UserGetResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2UserGetResult))
	} else if resType == "V2UserSetResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2UserSetResult))
	} else if resType == "V2ShortUrlCreateResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2ShortUrlCreateResult))
	} else if resType == "V2ShortUrlStatsResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2ShortUrlStatsResult))
	} else if resType == "V2VsmsTplBatchSendResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsTplBatchSendResult))
	} else if resType == "V2VsmsSignAddResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsSignAddResult))
	} else if resType == "V2VsmsSignSearchResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsSignSearchResult))
	} else if resType == "V2VsmsSignDeleteResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsSignDeleteResult))
	} else if resType == "V2VsmsTplAddResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsTplAddResult))
	} else if resType == "V2VsmsTplGetResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsTplGetResult))
	} else if resType == "V2VsmsTplDeleteResult" {
		err = json.Unmarshal([]byte(data), &((*r).V2VsmsTplDeleteResult))
	}
	return err
}

type Result struct {
	Status int //0 接口请求成功、-1 接口请求失败
	FailResult
	V2SingleSendResult
	V2BatchSendResult
	V2TplSingleSendResult
	V2TplBatchSendResult
	V2PullStatusResult
	V2PullReplyResult
	V2TplAddResult
	V2TplGetResult
	V2TplUpdateResult
	V2TplDeleteResult
	V2SignAddResult
	V2SignGetResult
	V2SignUpdateResult
	V2GetRecordResult
	V2RegCompleteResult
	V2GetTotalFeeResult
	V2VoiceSendResult
	V2VoicePullStatusResult
	V2UserGetResult
	V2UserSetResult
	V2ShortUrlCreateResult
	V2ShortUrlStatsResult
	V2VsmsTplBatchSendResult
	V2VsmsSignAddResult
	V2VsmsSignSearchResult
	V2VsmsSignDeleteResult
	V2VsmsTplAddResult
	V2VsmsTplGetResult
	V2VsmsTplDeleteResult
}

type FailResult struct {
	HttpStatusCode int    `json:"http_status_code"`
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	Detail         string `json:"detail"`
}

type V2SingleSendResult struct {
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Count  int     `json:"count"`
	Fee    float32 `json:"fee"`
	Unit   string  `json:"unit"`
	Mobile string  `json:"mobile"`
	Sid    int64   `json:"sid"`
}

type V2BatchSendResult struct {
	TotalCount int                  `json:"total_count"`
	TotalFee   string               `json:"total_fee"`
	Unit       string               `json:"unit"`
	Data       []V2SingleSendResult `json:"data"`
}

type V2TplSingleSendResult struct {
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Count  int     `json:"count"`
	Fee    float32 `json:"fee"`
	Unit   string  `json:"unit"`
	Mobile string  `json:"mobile"`
	Sid    int64   `json:"sid"`
}

type V2TplBatchSendResult struct {
	TotalCount int                  `json:"total_count"`
	TotalFee   string               `json:"total_fee"`
	Unit       string               `json:"unit"`
	Data       []V2SingleSendResult `json:"data"`
}

type V2Status struct {
	ErrorDetail     string `json:"error_detail"`
	Sid             int    `json:"sid"`
	Uid             int    `json:"uid"`
	UserReceiveTime string `json:"user_receive_time"`
	ErrorMsg        string `json:"error_msg"`
	ReportStatus    string `json:"report_status"`
}
type V2PullStatusResult []V2Status

type V2Reply struct {
	Id         string `json:"id"`
	Mobile     string `json:"mobile"`
	Text       string `json:"text"`
	ReplyTime  string `json:"reply_time"`
	Extend     string `json:"extend"`
	BaseExtend string `json:"base_extend"`
	Sign       string `json:"_sign"`
}

type V2PullReplyResult []V2Reply

type V2TplAddResult struct {
	TplId       int    `json:"tpl_id"`
	TplContent  string `json:"tpl_content"`
	CheckStatus string `json:"check_status"`
	Reason      string `json:"reason"`
}

type V2TplGetResult struct {
	TplId       int    `json:"tpl_id"`
	TplContent  string `json:"tpl_content"`
	CheckStatus string `json:"check_status"`
	Reason      string `json:"reason"`
	CountryCode string `json:"country_code"`
}

type V2TplUpdateResult struct {
	TplId       int    `json:"tpl_id"`
	TplContent  string `json:"tpl_content"`
	CheckStatus string `json:"check_status"`
	Reason      string `json:"reason"`
}

type V2TplDeleteResult struct {
	TplId       int    `json:"tpl_id"`
	TplContent  string `json:"tpl_content"`
	CheckStatus string `json:"check_status"`
	Reason      string `json:"reason"`
	CountryCode string `json:"country_code"`
}

type V2SignAdd struct {
	ApplyState   string `json:"apply_state"`
	Sign         string `json:"sign"`
	IsApplyVip   bool   `json:"is_apply_vip"`
	IsOnlyGlobal bool   `json:"is_only_global"`
	IndustryType string `json:"industry_type"`
}
type V2SignAddResult struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Sign V2SignAdd `json:"sign"`
}

type V2SignGet struct {
	Chan         string `json:"chan"`
	CheckStatus  string `json:"check_status"`
	Enabled      bool   `json:"enabled"`
	Extend       string `json:"extend"`
	IndustryType string `json:"industry_type"`
	OnlyGlobal   bool   `json:"only_global"`
	Remark       string `json:"remark"`
	Sign         string `json:"sign"`
	Vip          bool   `json:"vip"`
}

type V2SignGetResult struct {
	Code  int         `json:"code"`
	Total int         `json:"total"`
	Sign  []V2SignGet `json:"sign"`
}

type V2SignUpdateResult struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Sign V2SignAdd `json:"sign"`
}

type Record struct {
	Sid             int64  `json:"sid"`
	Mobile          string `json:"mobile"`
	SendTime        string `json:"send_time"`
	Text            string `json:"text"`
	SendStatus      string `json:"send_status"`
	ReportStatus    string `json:"report_status"`
	Fee             int    `json:"fee"`
	UserReceiveTime string `json:"user_receive_time"`
	Uid             string `json:"uid"`
}
type V2GetRecordResult []Record

type V2RegCompleteResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type V2GetTotalFeeResult struct {
	TotalCount        int    `json:"totalCount"`
	TotalFee          string `json:"totalFee"`
	TotalSuccessCount int    `json:"totalSuccessCount"`
	TotalFailedCount  int    `json:"totalFailedCount"`
	TotalUnknownCount int    `json:"totalUnknownCount"`
}

type V2VoiceSendResult struct {
	Count int     `json:"count"`
	Fee   float64 `json:"fee"`
	Sid   string  `json:"sid"`
}

type V2VoiceStatus struct {
	Sid             string `json:"sid"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	UserReceiveTime string `json:"user_receive_time"`
	Uid             string `json:"uid"`
	ErrorMsg        string `json:"error_msg"`
	Duration        int    `json:"duration"`
}

type V2VoicePullStatusResult []V2VoiceStatus

type V2UserGetResult struct {
	Nick             string  `json:"nick"`
	GmtCreated       string  `json:"gmt_created"`
	Mobile           string  `json:"mobile"`
	Email            string  `json:"email"`
	IpWhitelist      string  `json:"ip_whitelist"`
	ApiVersion       string  `json:"api_version"`
	Balance          float64 `json:"balance"`
	AlarmBalance     float64 `json:"alarm_balance"`
	EmergencyContact string  `json:"emergency_contact"`
	EmergencyMobile  string  `json:"emergency_mobile"`
}

type V2UserSetResult struct {
	Nick             string  `json:"nick"`
	GmtCreated       string  `json:"gmt_created"`
	Mobile           string  `json:"mobile"`
	Email            string  `json:"email"`
	IpWhitelist      string  `json:"ip_whitelist"`
	ApiVersion       string  `json:"api_version"`
	Balance          float64 `json:"balance"`
	AlarmBalance     float64 `json:"alarm_balance"`
	EmergencyContact string  `json:"emergency_contact"`
	EmergencyMobile  string  `json:"emergency_mobile"`
}

type ShortUrlT struct {
	Sid      string `json:"sid"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	EnterUrl string `json:"enter_url"`
	Name     string `json:"name"`
}

type V2ShortUrlCreateResult struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	ShortUrl ShortUrlT `json:"short_url"`
}

type V2ShortUrlRecord struct {
	Pt string `json:"pt"`
	Pv int    `json:"pv"`
	Uv int    `json:"uv"`
	Ip int    `json:"ip"`
}

type V2ShortUrlStatsResult struct {
	Records  []V2ShortUrlRecord `json:"records"`
	Size     int                `json:"size"`
	TimeStep string             `json:"time_step"`
}

type V2VsmsSend struct {
	Msg    string  `json:"msg"`
	Code   int     `json:"code"`
	Fee    float64 `json:"fee"`
	Mobile string  `json:"mobile"`
	Sid    int     `json:"sid"`
}

type V2VsmsTplBatchSendResult struct {
	TotalCount int          `json:"total_count"`
	TotalFee   string       `json:"total_fee"`
	Code       int          `json:"code"`
	Msg        string       `json:"msg"`
	Count      int          `json:"count"`
	Fee        float32      `json:"fee"`
	Unit       string       `json:"unit"`
	Mobile     string       `json:"mobile"`
	Sid        int64        `json:"sid"`
	Data       []V2VsmsSend `json:"data"`
}

type V2VsmsSignAddData struct {
	Sign string `json:"sign"`
}

type V2VsmsSignAddResult struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
	Data   string `json:"data"`
}

type V2VsmsSignSearchListEle struct {
	Sign   string `json:"sign"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

type V2VsmsSignSearchData struct {
	SignList []V2VsmsSignSearchListEle `json:"sign_list"`
	Total    int                       `json:"total"`
}

type V2VsmsSignSearchResult struct {
	HttpStatusCode int    `json:"http_status_code"`
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	Detail         string `json:"detail"`
	Data           string `json:"data"`
}

type V2VsmsSignDeleteResult struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
	Data   string `json:"data"`
}

type V2VsmsTplAddData struct {
	TplId int `json:"tpl_id"`
}

type V2VsmsTplAddResult struct {
	Code   int              `json:"code"`
	Msg    string           `json:"msg"`
	Detail string           `json:"detail"`
	Data   V2VsmsTplAddData `json:"data"`
}

type V2VsmsTplGetData struct {
	TplId         int    `json:"tpl_id"`
	CheckStatus   string `json:"check_status"`
	Reason        string `json:"reason"`
	YdEnabled     bool   `json:"yd_enabled"`
	LtEnabled     bool   `json:"lt_enabled"`
	DxEnables     bool   `json:"dx_enables"`
	YdCheckStatus int    `json:"yd_check_status"`
	LtCheckStatus int    `json:"lt_check_status"`
	DxCheckStatus int    `json:"dx_check_status"`
	YdCheckRemark string `json:"yd_check_remark"`
	LtCheckRemark string `json:"lt_check_remark"`
	DxCheckRemark string `json:"dx_check_remark"`
	ExpireTime    string `json:"expire_time"`
}

type V2VsmsTplGetResult struct {
	Code           int              `json:"code"`
	Msg            string           `json:"msg"`
	Detail         string           `json:"detail"`
	HttpStatusCode int              `json:"http_status_code"`
	Data           V2VsmsTplGetData `json:"data"`
}

type V2VsmsTplDeleteResult struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
	Data   string `json:"data"`
}
