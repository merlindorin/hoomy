package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/merlindorin/hoomy/cmd/filter"
	"github.com/merlindorin/hoomy/internal/cmd"
	"github.com/merlindorin/hoomy/pkg/kizbox"
	"go.uber.org/zap"
)

type DiscoverCmd struct {
	filter.Filter

	Timeout time.Duration `default:"5s" help:"timeout for discovering (ns, ms, s & m)"`
}

func (d *DiscoverCmd) Run(common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	logger = logger.With(zap.Duration("timeout", d.Timeout))
	ctx := context.Background()

	serverChan := make(chan *kizbox.Server, 4)
	endChan := make(chan error, 4)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		endChan <- kizbox.NewDiscover(kizbox.WithDisableIPv6(), kizbox.WithTimeout(d.Timeout)).Run(ctx, serverChan)
	}()

	go func() {
		for server := range serverChan {
			logger.Info("new server discovered", zap.Any("server", server))
		}
	}()

	select {
	case <-c:
		return nil
	case err = <-endChan:
		return err
	}
}
