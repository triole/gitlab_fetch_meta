package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/triole/logseal"
)

func fetchReposMeta(url string) (repos tReposMeta, err error) {
	by, err := req(url)
	lg.Trace("got data", logseal.F{"data": fmt.Sprintf("%s", by)})
	lg.IfErrError(err)
	if err == nil {
		err = json.Unmarshal(by, &repos)
		lg.IfErrError("can not unmarshal response", err)
	}
	return
}

func req(targetURL string) (data []byte, err error) {
	lg.Debug("fire request", logseal.F{"url": targetURL})
	url, err := url.Parse(targetURL)
	lg.IfErrError("can not parse url", logseal.F{"error": err})

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Set("User-Agent", CLI.UA)
	lg.IfErrError("can not init request", logseal.F{"error": err})

	//calling the URL
	response, err := client.Do(request)
	lg.IfErrError("request failed", logseal.F{"error": err})

	//getting the response
	data, err = io.ReadAll(response.Body)
	lg.IfErrError("unable to read request response", logseal.F{"error": err})
	if err != nil {
		return
	}
	return
}
