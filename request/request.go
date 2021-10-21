package request

import (
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/result"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

func HttpPostForm(apiUrl string, param param.RequestParam, resType string, apikey string) (result.Result, error){
	apiResult := result.Result{}
	var data = url.Values{}
	data.Set("apikey", apikey)
	refV := reflect.ValueOf(param)
	for i:=0; i<refV.NumField(); i++ {
		field := refV.Type().Field(i)
		tag := field.Tag
		paramFormName := tag.Get("param")
		if refV.Field(i).String() != "" || paramFormName == "tpl_value" {
			data.Set(paramFormName, refV.Field(i).String())
		}
	}
	fmt.Println(data)
	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return apiResult, err
	}
	if resp.StatusCode == 200 {
		apiResult.Status = 0
	} else {
		apiResult.Status = -1
		resType = "FailResult"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apiResult, err
	}
	fmt.Println(string(body))
	err = result.Format(string(body), &apiResult, resType)
	if err != nil {
		return apiResult, err
	}
	return apiResult, nil
}

func HttpGet(apiUrl string, param param.RequestParam, resType string, apikey string) (result.Result, error){
	apiResult := result.Result{}
	var data = "apikey=" + apikey + "&"
	refV := reflect.ValueOf(param)
	for i:=0; i<refV.NumField(); i++ {
		field := refV.Type().Field(i)
		tag := field.Tag
		paramFormName := tag.Get("param")
		if refV.Field(i).String() != "" || paramFormName == "tpl_value" {
			data += paramFormName + "=" + url.QueryEscape(refV.Field(i).String()) + "&"
		}
	}
	resp, err := http.Get(apiUrl + "?" + data)
	if err != nil {
		fmt.Println(err)
		return apiResult, err
	}
	if resp.StatusCode == 200 {
		apiResult.Status = 0
	} else {
		apiResult.Status = -1
		resType = "FailResult"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apiResult, err
	}
	fmt.Println(string(body), apiUrl + "?" + data)
	err = result.Format(string(body), &apiResult, resType)
	if err != nil {
		return apiResult, err
	}
	return apiResult, nil
}
