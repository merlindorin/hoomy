package commands

import (
	"context"
	"fmt"

	"github.com/merlindorin/hoomy/cmd/hoomy/filter"
	"github.com/merlindorin/hoomy/cmd/hoomy/globals"
	"github.com/merlindorin/hoomy/internal/cmd"
	v1 "github.com/merlindorin/hoomy/pkg/kizbox/api/v1"
)

type DevicesCmd struct {
	filter.Filter
}

func (d *DevicesCmd) Run(global *globals.Globals, common *cmd.Commons) error {
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
		logger.Info(fmt.Sprintf("Device %s (%s)", device.Label, device.DeviceURL))
	}

	return nil
}
