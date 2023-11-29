package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Printf("Error dumping request: %s\n", err)
	}
	fmt.Printf("REQUEST:\n%s\n\n", string(reqDump))
}

func PrintResponse(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("Error dumping respoonse: %s\n", err)
	}
	fmt.Printf("RESPONSE:\n%s\n\n", string(respDump))
}
