package request

import (
	"bytes"
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/param"
	"github.com/fixJ/yunpian-go-sdk/result"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	err = result.Format(string(body), &apiResult, resType)
	if err != nil {
		return apiResult, err
	}
	return apiResult, nil
}


func HttpPostMultiPartForm(apiUrl string, param param.RequestParam, resType string, apikey string, filePath string) (result.Result, error) {
	apiResult := result.Result{}
	file, err := os.Open(filePath)
	if err != nil {
		return apiResult, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("material", filepath.Base(filePath))
	if err != nil {
		return apiResult, err
	}
	_, err = io.Copy(part, file)
	refV := reflect.ValueOf(param)
	for i := 0; i < refV.NumField(); i++ {
		field := refV.Type().Field(i)
		tag := field.Tag
		paramFormName := tag.Get("param")
		if refV.Field(i).String() != "" || paramFormName == "tpl_value" {
			//data.Set(paramFormName, refV.Field(i).String())
			writer.WriteField(paramFormName, refV.Field(i).String())
		}
	}
	writer.WriteField("apikey", apikey)
	err = writer.Close()
	if err != nil {
		return apiResult, err
	}

	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		return apiResult, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode == 200 {
		apiResult.Status = 0
	} else {
		apiResult.Status = -1
		resType = "FailResult"
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apiResult, err
	}
	err = result.Format(string(respBody), &apiResult, resType)
	if err != nil {
		return apiResult, err
	}
	return apiResult, nil
}