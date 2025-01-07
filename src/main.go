package main

import (
	"encoding/json"
	"fmt"

	"github.com/triole/logseal"
)

var (
	conf tConf
)

func main() {
	parseArgs()
	conf = initConf(
		CLI.GitlabEP, CLI.MetaFilenames, CLI.UA,
		CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON,
	)
	arr, _ := fetchRepos()
	fmt.Printf("%+v\n", arr)
}

func fetchRepos() (repos tReposMeta, err error) {
	by, err := req(conf.GitlabEP)
	conf.Lg.Trace("got data", logseal.F{"data": fmt.Sprintf("%s", by)})
	conf.Lg.IfErrError(err)
	if err == nil {
		err = json.Unmarshal(by, &repos)
		conf.Lg.IfErrError("can not unmarshal response", err)
	}
	for _, el := range repos {
		el.MetaFilenames = conf.MetaFilenames
		el.initMeta()
	}
	return
}
