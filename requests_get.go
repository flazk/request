package request

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
)

type Submit struct {
	Url string               //请求地址
	Data map[string]string				 //提交得数据表单，用于POST请求
	Header map[string]string //请求头
	Cookies map[string]string //Cookie
	Proxy string             //代理IP

}

func(t *Submit)GET()(*http.Response,error){
	client:=&http.Client{}
	if t.Url == ""{
		//return a,errors.New()
		panic("url错误")
	}
	if t.Proxy != ""{
		proxyUrl, _ := url.Parse(t.Proxy)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	req,_:=http.NewRequest("GET",t.Url,nil)
	for key,value := range t.Header{
		req.Header.Set(key,value)
	}
	response,_:=client.Do(req)
	return response,nil
}

//识别响应内容编码
/*
bodyReader := bufio.NewReader(r.Body)
e := determineEncoding(bodyReader)
fmt.Printf("Encoding %v\n", e)
decodeReader := transform.NewReader(bodyReader, e.NewDecoder())
 */
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("err %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


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