package cmd

import (
	"github.com/urfave/cli/v2"
	"runtime/debug"
	"time"
)

const ourPath = "github.com/cita-cloud/executor_noop_go"
const VERSION = "v6.7.3"

type GitInfo struct {
	Commit string // head commit hash
	Date   string // commit time in YYYYMMDD format
	Dirty  bool
}

func NewGI() (GitInfo, bool) {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if buildInfo.Main.Path == ourPath {
			return buildInfoGit(buildInfo)
		}
	}
	return GitInfo{}, false
}

const (
	govcsTimeLayout = "2006-01-02T15:04:05Z"
	ourTimeLayout   = "20060102"
)

func buildInfoGit(info *debug.BuildInfo) (s GitInfo, ok bool) {
	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.revision":
			s.Commit = v.Value
		case "vcs.modified":
			if v.Value == "true" {
				s.Dirty = true
			}
		case "vcs.time":
			t, err := time.Parse(govcsTimeLayout, v.Value)
			if err == nil {
				s.Date = t.Format(ourTimeLayout)
			}
		}
	}
	if s.Commit != "" && s.Date != "" {
		ok = true
	}
	return
}

// NewApp creates an app with sane defaults.
func NewApp(usage string) *cli.App {
	git, _ := NewGI()
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = VersionWithCommit(git)
	app.Usage = usage
	app.Copyright = "Copyright 2024 The Cita-Cloud Authors"
	return app
}

func VersionWithCommit(git GitInfo) string {
	vsn := VERSION
	if len(git.Commit) >= 8 {
		vsn += "-" + git.Commit[:8]
	}
	if git.Date != "" {
		vsn += "-" + git.Date
	}
	if git.Dirty == true {
		vsn += "-dirty"
	}
	return vsn
}
