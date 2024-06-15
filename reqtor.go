package reqtor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"time"
	"runtime"
	"syscall"
)

var (
	ProxyProtocol = "socks5"
	ProxyHost     = "127.0.0.1"
	ProxyPort     = 9050
	AutoStart     = false
	AutoStop      = false
	Debugger      = false
)

var torCmd *exec.Cmd

func Start() error {
	if Debugger {
		logger(false, false, "Start S1", "TOR Starting...")
	}
	torCmd = exec.Command("tor")
	if runtime.GOOS == "windows"{
		torCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	err := torCmd.Start()
	if err != nil {
		logger(true, false, "Start E1", err.Error())
		return err
	}
	<-time.After(4 * time.Second)

	if Debugger {
		logger(false, false, "Start S2", "TOR Successfully Started")
	}
	return nil
}

func Stop() error {
	if Debugger {
		logger(false, false, "Stop S1", "TOR Stopping...")
	}
	err := torCmd.Process.Kill()
	if err != nil {
		if Debugger {
			logger(true, false, "Stop E1", "Failed to stop TOR ERROR:"+err.Error())
		}
		return err
	}
	if Debugger {
		logger(false, false, "Stop S2", "TOR Successfully Stoped")
	}
	return nil
}

func Request(requestType string, requestUrl string, requestHeaders map[string]string, requestBody map[string]interface{}) (*http.Response, error) {
	var startTime time.Time
	if AutoStart {
		Start()
	}

	if Debugger {
		startTime = time.Now()
		logger(false, false, "Request S1", "TOR Request Starting...")
	}

	var torProxy = fmt.Sprintf("%s://%s:%d", ProxyProtocol, ProxyHost, ProxyPort)
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		if Debugger {
			logger(true, false, "Request E1", err.Error())
		}
		return nil, err
	}

	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}

	//This function works, if the request type POST and the request body is filled
	reqBodyJSON := bytes.NewBuffer(nil)
	if requestType == "POST" && requestBody != nil {
		reqJson, err := json.Marshal(requestBody)
		if err != nil {
			if Debugger {
				logger(true, false, "Request E2", err.Error())
			}
			return nil, err
		}
		reqBodyJSON = bytes.NewBuffer(reqJson)
	}
	req, err := http.NewRequest(requestType, requestUrl, reqBodyJSON)
	if err != nil {
		if Debugger {
			logger(true, false, "Request E3", err.Error())
		}
		return nil, err
	}

	if reqBodyJSON != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range requestHeaders {
		if Debugger {
			logger(false, false, "Request E4", "Setting The Request Headers... "+k)
		}
		req.Header.Set(k, v)
	}

	client := &http.Client{Transport: torTransport}
	response, err := client.Do(req)
	if err != nil {
		if Debugger {
			logger(true, false, "Request E5", err.Error())
		}
		return nil, err
	}

	if AutoStop {
		err := Stop()
		if err != nil {
			if Debugger {
				logger(true, false, "Request E6", err.Error())
			}
			return nil, err
		}
	}

	if Debugger {
		duration := time.Since(startTime).Seconds()
		logger(false, false, "Request S2", fmt.Sprintf("Request Successfully. Duration: %.2f seconds", duration))
	}
	return response, nil
}

// GET - Request
func Get(requestUrl string, requestHeaders map[string]string) (*http.Response, error) {
	response, err := Request("GET", requestUrl, requestHeaders, nil)
	if err != nil {
		if Debugger {
			logger(true, false, "Get E1", err.Error())
		}
		return nil, err
	}
	if Debugger {
		logger(false, false, "Get S1", "GET Request Successfully")
	}
	return response, nil

}

// POST - Request
func Post(requestUrl string, requestHeaders map[string]string, requestBody map[string]interface{}) (*http.Response, error) {
	response, err := Request("POST", requestUrl, requestHeaders, requestBody)
	if err != nil {
		if Debugger {
			logger(true, false, "Post E1", err.Error())
		}
		return nil, err
	}
	if Debugger {
		logger(false, false, "Post S1", "POST Request Successfully")
	}
	return response, nil
}

// This Function Returns the GET request body
func GetBody(requestUrl string, requestHeaders map[string]string) ([]byte, error) {
	response, err := Request("GET", requestUrl, requestHeaders, nil)
	if err != nil {
		if Debugger {
			logger(true, false, "GetBody E1", err.Error())
		}
		return nil, err
	}
	defer response.Body.Close()

	body, err := ResBody(response.Body)
	if err != nil {
		if Debugger {
			logger(true, false, "GetBody E2", err.Error())
		}
	}
	if Debugger {
		logger(false, false, "GetBody S1", "GetBody Successfully")
	}
	return body, nil
}

// This Function Returns the POST request body
func PostBody(requestUrl string, requestHeaders map[string]string, requestBody map[string]interface{}) ([]byte, error) {
	response, err := Request("POST", requestUrl, requestHeaders, requestBody)
	if err != nil {
		if Debugger {
			logger(true, false, "PostBody E1", err.Error())
		}
		return nil, err
	}
	defer response.Body.Close()

	body, err := ResBody(response.Body)
	if err != nil {
		if Debugger {
			logger(true, false, "PostBody E2", err.Error())
		}
	}
	if Debugger {
		logger(false, false, "PostBody S1", "PostBody Successfully")
	}
	return body, nil
}

// This Function Returns the HTTP Response Body
func ResBody(response io.Reader) ([]byte, error) {
	body, err := io.ReadAll(response)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func logger(isError bool, isTest bool, runFunc string, body string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	command := "\033[32m"
	if isError {
		command = "\033[31m"
	}
	header := "Reqtor"
	if isTest {
		header += " TEST"
	}
	command += "[ " + timestamp + " ]-[ " + header + " ]--[ " + runFunc + " ]--[ " + body + " ]\033[0m"
	fmt.Println(command)
}
