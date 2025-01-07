package main

import (
	"fmt"

	"github.com/triole/logseal"
)

var (
	lg = logseal.Init()
)

func main() {
	parseArgs()
	lg = logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	arr, _ := fetchReposMeta(CLI.GitlabEP)
	fmt.Printf("%+v\n", arr)
}
