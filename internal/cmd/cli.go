package cmd

// An Option applies optional changes to the CLI application.
type Option func(runner *Runner) error

func (o Option) Apply(r *Runner) error {
	return o(r)
}

type Runner struct {
	name        string
	description string
	authors     []string
	version     string

	commands interface{}
}

func WithName(name string) Option {
	return func(r *Runner) error {
		r.name = name
		return nil
	}
}

func WithDescription(description string) Option {
	return func(r *Runner) error {
		r.description = description
		return nil
	}
}

func WithAuthors(authors ...string) Option {
	return func(r *Runner) error {
		r.authors = authors
		return nil
	}
}

func WithVersion(version string) Option {
	return func(r *Runner) error {
		r.version = version
		return nil
	}
}

// Read the doc
func WithCommands(commands interface{}) Option {
	return func(r *Runner) error {
		r.commands = commands
		return nil
	}
}

func New(opts ...Option) *Runner {
	return nil
}
