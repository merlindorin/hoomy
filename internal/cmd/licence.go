package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// Licence represent a license with its content. It is not possible to use another primitive (like string) because
// kong will not recognize the licence command
type Licence struct {
	content string
}

func NewLicence(s string) Licence {
	return Licence{content: s}
}

func (l Licence) Run(*kong.Context) (err error) {
	fmt.Printf("%s", l.content)
	return nil
}
