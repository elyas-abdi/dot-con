package con

import `fmt`

type Con struct {
	args    map[string]string
	context map[string]Detail
	dir     string
}

func New(options ...Option) (*Con, error) {
	c := Con{
		args:    make(map[string]string),
		context: make(map[string]Detail),
	}

	settings := Options{
		args: make(map[string]string),
	}

	err := c.resolveOptions(&settings, options...)
	if err != nil {
		return nil, err
	}

	c.args = settings.args
	c.dir = settings.dir

	err = c.parseDir()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", c)

	return &c, nil
}

func (c *Con) resolveOptions(settings *Options, options ...Option) error {
	for _, option := range options {
		err := option(settings)
		if err != nil {
			return err
		}
	}

	return nil
}
