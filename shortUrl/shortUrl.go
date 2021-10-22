package shortUrl

import (
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/request"
	"github.com/fixJ/yunpian-go-sdk/result"
)

type ShortUrl struct {
	Apikey string
	Domain string
}

// V2ShortUrlCreate 创建短链接
func (s ShortUrl) V2ShortUrlCreate(param param.V2ShortUrlCreateParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/short_url/shorten.json"
	return request.HttpPostForm(apiUrl, param, "V2ShortUrlCreateResult", s.Apikey)
}

// V2ShortUrlStats 短链接统计
func (s ShortUrl) V2ShortUrlStats(param param.V2ShortUrlStatsParam) (result.Result, error) {
	apiUrl := s.Domain + "/v2/short_url/time_series_stats.json"
	return request.HttpGet(apiUrl, param, "V2ShortUrlStatsResult", s.Apikey)
}