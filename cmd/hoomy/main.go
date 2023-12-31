package main

import (
	_ "embed"

	"github.com/alecthomas/kong"
	"github.com/merlindorin/hoomy/cmd/hoomy/commands"
	"github.com/merlindorin/hoomy/cmd/hoomy/globals"
	"github.com/merlindorin/hoomy/internal/cmd"
)

var (
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

type CLI struct {
	*cmd.Commons
	*globals.Globals

	Venitian commands.VenitianCmd `cmd:"venitians"`
	Devices  commands.DeviceCmd   `cmd:"devices" help:"list devices availables in the current system"`
	Listen   commands.ListenCmd   `cmd:"listen" help:"listen events in the current system"`
	Discover commands.DiscoverCmd `cmd:"discover" help:"list for systems available"`
}

func main() {
	cli := CLI{
		Commons: &cmd.Commons{
			Version: cmd.NewVersion(version, commit, buildSource, date),
			Licence: cmd.NewLicence(license),
		},
		Globals: &globals.Globals{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name("hoomy"),
		kong.Description("Simple cli for managing my home automation"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Globals, cli.Commons))
}
