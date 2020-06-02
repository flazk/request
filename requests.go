package requests

import (
	"net/http"
	"net/url"
)



func Requests(dict map[string]string,header map[string]string)(*http.Response,error){
	/*
		dict["url"]为传递的url地址
		dict["method"]为传递的请求方式，GET、POST
		dict["proxy"]为代理地址
		header[] 为传递请求头
	 */
	client:=&http.Client{}
	if dict["url"] == ""{
		//return a,errors.New()
		panic("url错误")
	}
	if dict["method"] == ""{
		dict["method"] = "GET"
	}
	if dict["proxy"] != ""{
		proxyUrl, _ := url.Parse(dict["proxy"])
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	req,_:=http.NewRequest(dict["method"],dict["url"],nil)
	for key,value := range header{
		req.Header.Set(key,value)
	}
	response,_:=client.Do(req)
	return response,nil
}