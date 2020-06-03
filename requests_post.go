package request

import (
	"net/http"
	"net/url"
	"request/requests_get"
	"strings"
)

func (t *Submit)POST()(*http.Response,error) {
	client := &http.Client{}
	if t.Url == "" {
		panic("url错误")
	}
	if t.Proxy != "" {
		proxyUrl, _ := url.Parse(t.Proxy)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	payload := make(url.Values)
	for key, value := range t.Data {
		payload.Add(key, value)
	}
	req, _ := http.NewRequest("POST",t.Url,strings.NewReader(payload.Encode()))
	for key, value := range t.Header {
		req.Header.Set(key, value)
	}
	for key, value := range t.Cookies {
		req.AddCookie(&http.Cookie{
			Name:key,
			Value:value,
		})
		response, _ := client.Do(req)
		return response, nil
	}
}
/*
req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies", nil)
if err != nil {
	panic(err)
}

req.AddCookie(&http.Cookie{
	Name:   "name",
	Value:  "poloxue",
	Domain: "httpbin.org",
	Path:   "/cookies",
})

r, err := http.DefaultClient.Do(req)
 */