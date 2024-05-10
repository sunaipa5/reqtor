
# Reqtor

[![Go Reference](https://pkg.go.dev/badge/github.com/sunaipa5/reqtor.svg)](https://pkg.go.dev/github.com/sunaipa5/reqtor)

## Prerequirement 
In order to use the AutoStart and AutoStop features, the ``tor`` cli tool is required and must be added to bash, AutoStart and AutoStop features are on by default.

## Sample Request

```
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
)


func main() {
	html, _ := reqtor.GetBody("https://example.com", nil)
	fmt.Println(string(html))
}

```
## Request with headers
```
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
	"io"
)

var RequestHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Encoding": "gzip,deflate,sdch",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
}

func main() {

	response, _ := reqtor.Get("https://example.com", RequestHeaders)

	html, _ := io.ReadAll(response.Body)

	reqtor.CloseResponse(response)

	fmt.Println(string(html))
}


```
### Manuel Usage
if you are launching tor yourself or using the port on the tor browser you can set it manually
```
package main

import (
	"fmt"
	"github.com/sunaipa5/reqtor"
	"io"
)


func main() {
	reqtor.AutoStart = false
	reqtor.AutoStop = false
	reqtor.TorPort = 9150 //Tor Browser Default Port Number

	response, _ := reqtor.Get("https://example.com", nil)

	html, _ := io.ReadAll(response.Body)

	reqtor.CloseResponse(response)

	fmt.Println(string(html))

}

```
