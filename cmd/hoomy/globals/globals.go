package globals

import (
	"fmt"

	"github.com/merlindorin/hoomy/pkg/kizbox/client"
)

type Globals struct {
	ApiKey string `env:"API_KEY" help:"apikey (retrieved through developer quickstart)"`
	Host   string `env:"HOST" help:"host of the Kizbox"`
	Port   int    `default:"8443" help:"port of the Kizbox"`
}

func (c *Globals) Client() *client.ApiClient {
	return client.NewClient(fmt.Sprintf("%s:%d", c.Host, c.Port), c.ApiKey)
}
