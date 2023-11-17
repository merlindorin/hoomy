package commands

import (
	"context"
	"fmt"
	"io"
	"slices"

	"github.com/merlindorin/hoomy/cmd/hoomy/filter"
	"github.com/merlindorin/hoomy/cmd/hoomy/globals"
	"github.com/merlindorin/hoomy/internal/cmd"
	v1 "github.com/merlindorin/hoomy/pkg/kizbox/api/v1"
)

type DeviceCmd struct {
	List DevicesListCmd `cmd:"list" help:"list devices available"`
	Get  DevicesGetCmd  `cmd:"get" help:"get device by URL"`
}

type DevicesListCmd struct {
	filter.Filter
}

func (d *DevicesListCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := global.Client()
	ctx := context.Background()

	var devices []v1.Device
	_, err = cl.V1.Devices.List(ctx, &devices)
	if err != nil {
		logger.Error("cannot list devices")
		return err
	}

	for _, device := range devices {
		if (len(d.URLS) == 0 && len(d.Labels) == 0) ||
			slices.Contains(d.URLS, device.DeviceURL) ||
			slices.Contains(d.Labels, device.Label) {
			logger.Info(fmt.Sprintf("Device %s (%s)", device.Label, device.DeviceURL))
		}
	}

	return nil
}

type DevicesGetCmd struct {
	URL string `arg:"URL"`
}

func (d *DevicesGetCmd) Run(global *globals.Globals, common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := global.Client()
	ctx := context.Background()

	res, err := cl.V1.Devices.Get(ctx, d.URL, nil)
	if err != nil {
		logger.Error("cannot get device")
		return err
	}

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s", all)

	return nil
}
