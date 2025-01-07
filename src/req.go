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
	conf.Lg.Trace("got data", logseal.F{"data": fmt.Sprintf("%s", by)})
	conf.Lg.IfErrError(err)
	if err == nil {
		err = json.Unmarshal(by, &repos)
		conf.Lg.IfErrError("can not unmarshal response", err)
	}
	return
}

func req(targetURL string) (data []byte, err error) {
	conf.Lg.Debug("fire request", logseal.F{"url": targetURL})
	url, err := url.Parse(targetURL)
	conf.Lg.IfErrError("can not parse url", logseal.F{"error": err})

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Set("User-Agent", CLI.UA)
	conf.Lg.IfErrError("can not init request", logseal.F{"error": err})

	//calling the URL
	response, err := client.Do(request)
	conf.Lg.IfErrError("request failed", logseal.F{"error": err})

	//getting the response
	data, err = io.ReadAll(response.Body)
	conf.Lg.IfErrError("unable to read request response", logseal.F{"error": err})
	if err != nil {
		return
	}
	return
}
