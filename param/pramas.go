package param

type RequestParam interface {
}

type V2SingleSendParam struct {
	Mobile      string `param:"mobile"`
	Text        string `param:"text"`
	Extend      string `param:"extend"`
	Uid         string `param:"uid"`
	CallbackUrl string `param:"callback_url"`
	Register    string `param:"register"`
	MobileStat  string `param:"mobile_stat"`
}

type V2BatchSendParam struct {
	Mobile      string `param:"mobile"`
	Text        string `param:"text"`
	Extend      string `param:"extend"`
	Uid         string `param:"uid"`
	CallbackUrl string `param:"callback_url"`
	MobileStat  string `param:"mobile_stat"`
}

type V2TplSingleSendParam struct {
	Mobile      string `param:"mobile"`
	TplId       string `param:"tpl_id"`
	TplValue    string `param:"tpl_value"`
	CallbackUrl string `param:"callback_url"`
	Extend      string `param:"extend"`
	Uid         string `param:"uid"`
}

type V2TplBatchSendParam struct {
	Mobile      string `param:"mobile"`
	TplId       string `param:"tpl_id"`
	TplValue    string `param:"tpl_value"`
	CallbackUrl string `param:"callback_url"`
	Extend      string `param:"extend"`
	Uid         string `param:"uid"`
}

type V2PullStatusParam struct {
	PageSize string `param:"page_size"`
}

type V2PullReplyParam struct {
	PageSize string `param:"page_size"`
}

type V2TplAddParam struct {
	TplContent       string `param:"tpl_content"`
	NotifyType       string `param:"notify_type"`
	Website          string `param:"website"`
	TplType          string `param:"tplType"`
	Callback         string `param:"callback"`
	ApplyDescription string `param:"apply_description"`
	Region           string `param:"region"`
}

type V2TplGetParam struct {
	TplId string `param:"tpl_id"`
}

type V2TplUpdateParam struct {
	TplId            string `param:"tpl_id"`
	TplContent       string `param:"tpl_content"`
	NotifyType       string `param:"notify_type"`
	Website          string `param:"website"`
	TplType          string `param:"tplType"`
	Callback         string `param:"callback"`
	ApplyDescription string `param:"apply_description"`
	Region           string `param:"region"`
}

type V2TplDeleteParam struct {
	TplId string `param:"tpl_id"`
}

type V2SignAddParam struct {
	Sign          string `param:"sign"`
	NotifyType    string `param:"notify_type"`
	Website       string `param:"website"`
	ApplyVip      string `param:"apply_vip"`
	IsOnlyGlobal  string `param:"is_only_global"`
	IndustryType  string `param:"industry_type"`
	ProveType     string `param:"prove_type"`
	LicenseUrl    string `param:"license_url"`
	LicenseBase64 string `param:"license_base64"`
}

type V2SignGetParam struct {
	Sign     string `param:"sign"`
	PageNum  string `param:"page_num"`
	PageSize string `param:"page_size"`
}

type V2SignUpdateParam struct {
	OldSign       string `param:"old_sign"`
	Sign          string `param:"sign"`
	NotifyType    string `param:"notify_type"`
	Website       string `param:"website"`
	ApplyVip      string `param:"apply_vip"`
	IsOnlyGlobal  string `param:"is_only_global"`
	IndustryType  string `param:"industry_type"`
	LicenseUrl    string `param:"license_url"`
	LicenseBase64 string `param:"license_base64"`
}

type V2GetRecordParam struct {
	Mobile    string `param:"mobile"`
	StartTime string `param:"start_time"`
	EndTime   string `param:"end_time"`
	PageNum   string `param:"page_num"`
	PageSize  string `param:"page_size"`
	Type      string `param:"type"`
}

type V2RegCompleteParam struct {
	Mobile string `param:"mobile"`
	Time   string `param:"time"`
}

type V2GetTotalFeeParam struct {
	Date string `param:"date"`
}

type V2VoiceSendParam struct {
	Mobile      string `param:"mobile"`
	Code        string `param:"code"`
	Language    string `param:"language"`
	Uid         string `param:"uid"`
	CallbackUrl string `param:"callback_url"`
}

type V2VoicePullStatusParam struct {
	PageSize string `param:"page_size"`
}

type V2UserGetParam struct {
}

type V2UserSetParam struct {
	AlarmBalance     string `param:"alarm_balance"`
	EmergencyContact string `param:"emergency_contact"`
	EmergencyMobile  string `param:"emergency_mobile"`
}

type V2ShortUrlCreateParam struct {
	LongUrl string `param:"long_url"`
	Name    string `param:"name"`
}

type V2ShortUrlStatsParam struct {
	Sid       string `param:"sid"`
	StartTime string `param:"start_time"`
	EndTime   string `param:"end_time"`
}
