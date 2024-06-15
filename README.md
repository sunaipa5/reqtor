
# Reqtor
[![Go Reference](https://pkg.go.dev/badge/github.com/sunaipa5/reqtor.svg)](https://pkg.go.dev/github.com/sunaipa5/reqtor)

Go library for sending http requests through Tor proxy or any proxy

## Requirements
In order to use the Start, Stop Functions and AutoStart, AutoStop features, the `tor` cli tool must be added to the environment variables, the proxy is set to `socks5://127.0.0.1:9050` by default, you can change the proxy settings as you wish. 

## Examples
### Sample Get Request

```go
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
)


func main() {
	err := reqtor.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	resBody, err := reqtor.GetBody("https://example.com", nil)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(resBody))

	err = reqtor.Stop()
	if err != nil {
		fmt.Println(err)
	}
}

```
### Send requests with headers and manually make get request
Note : Headers also work with `GetBody` and `PostBody` Functions

```go
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
	"io"
)

var RequestHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
}

func main() {
    err := reqtor.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := reqtor.Get("https://example.com", RequestHeaders)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	
	resBody, err := io.ReadAll(response.Body)
	if err != nil{
		fmt.Println()
		return
	}
	
	fmt.Println(string(resBody))

    err = reqtor.Stop()
	if err != nil {
		fmt.Println(err)
	}
}


```
### Change Settings
If you want `Tor` to turn off and on automatically every time you make a request, you should set AutoStar and AutoStop to `true`, by default they are `false`. If you want to make requests through a different proxy you can change the `ProxyProtocol`, `ProxyHost` and `ProxyPort` settings, the default values are socks5://127.0.0.1:9050
```go
    package main

    import (
        "fmt"
        "github.com/sunaipa5/reqtor"
    )


    func main() {
        reqtor.ProxyHost = "socks5"    //Default Protocol - socks5
        reqtor.ProxyHost = "127.0.0.1" //Default Host - 127.0.0.1
        reqtor.ProxyPort = 9050        //Default Port - 9050

        reqtor.AutoStart = true //This setting starts automaticly tor cli tool before every request
        reqtor.AutoStop = true  //This setting stops automaticly tor cli tool after every request

        resBody, err := reqtor.GetBody("https://example.com", nil)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(string(resBody))
    }

```
### Post Request

```go
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
)

var RequestHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
	"Cookie":          "test_cookie=cookie_value; test_cookie2=cookie_value2",
}

var PostData = map[string]string{
	"name":     "john",
	"surename": "doe",
	"age":      "24",
}

func main() {
	err := reqtor.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	// If you are not sending RequestHeader or Post data, replace with nil
	resBody, err := reqtor.PostBody("https://example.com", RequestHeaders, PostData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resBody))

	err = reqtor.Stop()
	if err != nil {
		fmt.Println(err)
	}
}

```