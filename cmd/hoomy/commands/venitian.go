package commands

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/merlindorin/hoomy/cmd/hoomy/filter"
	"github.com/merlindorin/hoomy/cmd/hoomy/globals"
	"github.com/merlindorin/hoomy/internal/cmd"
	"github.com/merlindorin/hoomy/pkg/kizbox/api/v1"
	"github.com/merlindorin/hoomy/pkg/kizbox/client"
	"go.uber.org/zap"
)

const (
	ControllableName = "io:ExteriorVenetianBlindIOComponent"
)

type VenitianCmd struct {
	filter.Filter
	List  VenitianListCmd  `cmd:"list" help:"List stores. By default, it will list all stores"`
	Set   VenitianSetCmd   `cmd:"set" help:"Set stores. By default, it will set all stores"`
	Open  VenitianOpenCmd  `cmd:"open" help:"Open stores. By default, it will open all stores"`
	Close VenitianCloseCmd `cmd:"close" help:"Close stores. By default, it will close all stores"`
	Stop  VenitianStopCmd  `cmd:"stop" help:"Stop stores. By default, it will stop all stores"`
	Wink  VenitianWinkCmd  `cmd:"wink" help:"Wink stores. By default, it will wink all stores"`
	My    VenitianMyCmd    `cmd:"my" help:"Go to my position. By default, it will my all stores"`
}

type VenitianListCmd struct{}

func (s VenitianListCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	devices, res, err := DeviceList(ctx, api, []string{ControllableName}, parent.Filter)
	if err != nil {
		logger.Error("cannot list device", zap.Any("res", res))
		return err
	}

	for _, device := range devices {
		logger.Info(fmt.Sprintf("%s (%s)", device.Label, device.DeviceURL), zap.String("label", device.Label), zap.String("url", device.DeviceURL), zap.Bool("available", device.Available))
	}

	return nil
}

type VenitianOpenCmd struct{}

func (s VenitianOpenCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	_, err = DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, v1.Command{Name: "open"})
	if err != nil {
		return err
	}

	return nil
}

type VenitianSetCmd struct {
	Position             *int
	Orientation          *int
	Closure              *int
	Name                 *string
	MemorizedPosition    *int
	MemorizedOrientation *int
}

func (s VenitianSetCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	var commands []v1.Command

	// this sequence must be executed in this specific order
	if s.Orientation != nil {
		commands = append(commands, v1.Command{
			Name:       "setOrientation",
			Parameters: []interface{}{s.Orientation},
		})
	}

	if s.Position != nil {
		commands = append(commands, v1.Command{
			Name:       "setPosition",
			Parameters: []interface{}{s.Position},
		})
	}

	if s.Closure != nil {
		commands = append(commands, v1.Command{
			Name:       "setClosure",
			Parameters: []interface{}{s.Closure},
		})
	}

	if s.Name != nil {
		commands = append(commands, v1.Command{
			Name:       "setName",
			Parameters: []interface{}{s.Name},
		})
	}

	if s.MemorizedPosition != nil {
		commands = append(commands, v1.Command{
			Name:       "setMemorized1Position",
			Parameters: []interface{}{s.MemorizedPosition},
		})
	}

	if s.MemorizedOrientation != nil {
		commands = append(commands, v1.Command{
			Name:       "setMemorized1Orientation",
			Parameters: []interface{}{s.MemorizedPosition},
		})
	}

	commands = append(commands, v1.Command{
		Name:       "setMemorized1Position",
		Parameters: []interface{}{60},
	})

	commands = append(commands, v1.Command{
		Name:       "setMemorized1Orientation",
		Parameters: []interface{}{0},
	})

	res, err := DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, commands...)
	if err != nil {
		all, _ := io.ReadAll(res.Body)
		logger.Error("cannot dispatch", zap.Error(err), zap.Int("statusCode", res.StatusCode), zap.ByteString("body", all))
		return err
	}

	return nil
}

type VenitianMyCmd struct{}

func (s VenitianMyCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	_, err = DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, v1.Command{Name: "my"})
	if err != nil {
		return err
	}

	return nil
}

type VenitianWinkCmd struct{}

func (s VenitianWinkCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	_, err = DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, v1.Command{Name: "wink"})
	if err != nil {
		return err
	}

	return nil
}

type VenitianStopCmd struct{}

func (s VenitianStopCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	_, err = DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, v1.Command{Name: "stop"})
	if err != nil {
		return err
	}

	return nil
}

type VenitianCloseCmd struct{}

func (s VenitianCloseCmd) Run(global *globals.Globals, common *cmd.Commons, parent *VenitianCmd) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api := global.Client()

	_, err = DispatchDeviceAction(ctx, api, logger, []string{ControllableName}, parent.Filter, v1.Command{Name: "close"})
	if err != nil {
		return err
	}

	return nil
}

func DispatchDeviceAction(ctx context.Context, cl *client.ApiClient, logger *zap.Logger, controllers []string, filter filter.Filter, commands ...v1.Command) (*http.Response, error) {
	devices, res, err := DeviceList(ctx, cl, controllers, filter)
	if err != nil {
		logger.Error("cannot list device", zap.Any("res", res))
		return res, err
	}

	var actions []v1.Action
	for _, device := range devices {
		logger.Debug("open device", zap.Any("device", device))
		action := v1.Action{
			Commands:  commands,
			DeviceURL: device.DeviceURL,
		}

		actions = append(actions, action)
	}

	return cl.V1.Execution.Apply(ctx, v1.Execute{Label: "cli command test", Actions: actions}, nil)
}

func DeviceList(ctx context.Context, cl *client.ApiClient, controllers []string, filter filter.Filter) ([]v1.Device, *http.Response, error) {
	var allDevices []v1.Device
	res, err := cl.V1.Devices.List(ctx, &allDevices)
	if err != nil {
		return nil, res, fmt.Errorf("cannot get devices: %w", err)
	}

	var venitianDevices []v1.Device
	for _, device := range allDevices {
		if slices.Contains(controllers, device.ControllableName) {
			if (len(filter.URLS) == 0 && len(filter.Labels) == 0) ||
				slices.Contains(filter.URLS, device.DeviceURL) ||
				slices.Contains(filter.Labels, device.Label) {
				venitianDevices = append(venitianDevices, device)
			}
		}
	}

	return venitianDevices, res, nil
}
