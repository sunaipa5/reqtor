package reqtor

import (
	"fmt"
	"testing"
)

var RequestHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
	"Cookie":          "test_cookie=cookie_value; test_cookie2=cookie_value2",
}

var PostData = map[string]interface{}{
	"name":     "john",
	"surename": "doe",
	"age":      "24",
}

func Test_Start(t *testing.T) {
	Debugger = true
	err := Start()
	if err != nil {
		logger(true, true, "Start E1", "ERROR: "+err.Error())
	}
	logger(false, true, "Start S1", "Tor is Started")
}

func Test_GetBody(t *testing.T) {
	reqBody, err := GetBody("https://check.torproject.org/api/ip", nil)
	if err != nil {
		logger(true, true, "GetBody E1", "ERROR: "+err.Error())
		return
	}

	fmt.Println(string(reqBody))
	logger(false, true, "GetBody S1", "SUCCESSFULLY")
}

func Test_GetRequestWithHeaders(t *testing.T) {
	reqBody, err := GetBody("https://httpbin.org/headers", RequestHeaders)
	if err != nil {
		logger(true, true, "GetRequestWithHeaders E1", err.Error())
	}
	fmt.Println(string(reqBody))
}

func Test_PostRequest(t *testing.T) {
	resBody, err := PostBody("https://httpbin.org/post", nil, PostData)
	if err != nil {
		logger(true, true, "PostRequuest E1", err.Error())
	}
	logger(false, true, "PostRequest S1", "POST Request Sucessfully")
	fmt.Println(string(resBody))
}

func Test_PostRequestWithHeaders(t *testing.T) {
	resBody, err := PostBody("https://httpbin.org/post", RequestHeaders, PostData)
	if err != nil {
		logger(true, true, "PostRequuest E1", err.Error())
	}
	logger(false, true, "PostRequest S1", "POST Request Sucessfully")
	fmt.Println(string(resBody))
}

func Test_Stop(t *testing.T) {
	err := Stop()
	if err != nil {
		fmt.Println(err)
	}
	logger(false, true, "Stop S1", "Tor is Closed")
}
