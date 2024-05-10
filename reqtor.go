package reqtor

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

var (
	TorPort   = 9050
	AutoStart = true
	AutoStop  = true
)

var torCmd *exec.Cmd

func Start() error {
	torCmd = exec.Command("tor")
	err := torCmd.Start()
	if err != nil {
		return err
	}
	<-time.After(4 * time.Second)

	return nil
}

func Stop() error {
	err := torCmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

func Request(requestType string, requestUrl string, requestHeaders map[string]string) (*http.Response, error) {
	if AutoStart {
		Start()
	}
	var torProxy = fmt.Sprintf("socks5://127.0.0.1:%d", TorPort)
	fmt.Println(torProxy)
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		return nil, err
	}

	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}

	req, err := http.NewRequest(requestType, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if len(requestHeaders) == 0 {
		for k, v := range requestHeaders {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Transport: torTransport}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func CloseResponse(response *http.Response) error {

	defer response.Body.Close()

	if AutoStop {
		err := Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(requestUrl string, requestHeaders map[string]string) (*http.Response, error) {
	response, err := Request("GET", requestUrl, requestHeaders)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func Post(requestUrl string, requestHeaders map[string]string) (*http.Response, error) {
	response, err := Request("POST", requestUrl, requestHeaders)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetBody(requestUrl string, requestHeaders map[string]string) ([]byte, error) {
	response, err := Request("GET", requestUrl, requestHeaders)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	CloseResponse(response)
	return body, nil
}
