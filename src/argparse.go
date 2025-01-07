package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "gitlab_fetch_meta"
	appDescription = "fetch repo metadata from gitlab api endpoint"
	appMainversion = "0.1"
)

var CLI struct {
	GitlabEP      string   `help:"file to process, required positional arg" arg:"" optional:""`
	MetaFilenames []string `help:"file names of additional metadata to fetch" short:"a" sep:"," default:"public.toml"`
	UA            string   `help:"requests user agent" short:"u" default:"Mozilla/5.0 (X11; Linux x86_64; rv:133.0) Gecko/20100101 Firefox/133.0"`
	LogFile       string   `help:"log file" default:"/dev/stdout"`
	LogLevel      string   `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors   bool     `help:"disable output colours, print plain text"`
	LogJSON       bool     `help:"enable json log, instead of text one"`
	VersionFlag   bool     `help:"display version" short:"V"`
}

func parseArgs() {
	curdir, _ := os.Getwd()
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"curdir": curdir,
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}

	if CLI.GitlabEP == "" {
		fmt.Printf(
			"[error] %s\n",
			"gitlab endpoint url required, i.e. https://gitlab.com/api/v4/groups/groupname/projects?include_subgroups=true")
		os.Exit(1)
	}

	// ctx.FatalIfErrorf(err)
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}

func alnum(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z0-9_-]")
	return re.ReplaceAllString(s, "-")
}
