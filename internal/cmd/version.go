package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type Version struct {
	version     string
	commit      string
	buildSource string
	date        string
}

func NewVersion(version string, commit string, buildSource string, date string) Version {
	return Version{version: version, commit: commit, buildSource: buildSource, date: date}
}

func (u Version) Run(*kong.Context) (err error) {
	fmt.Printf("version=%s, commit=%s, buildDate=%s, buildSource=%s", u.version, u.commit, u.date, u.buildSource)
	return
}
