package main

import "github.com/triole/logseal"

type tConf struct {
	GitlabEP      string
	MetaFilenames []string
	UA            string
	Lg            logseal.Logseal
}

func initConf(
	gitlabEP string, metaFilenames []string, ua string,
	logLevel, logFile string, logNoColors, logJSON bool,
) (conf tConf) {
	conf.GitlabEP = gitlabEP
	conf.MetaFilenames = metaFilenames
	conf.UA = ua
	conf.Lg = logseal.Init(logLevel, logFile, logNoColors, logJSON)
	return
}
