package con

type Con struct {
	args    map[string]string
	context map[string]detail
	dir     string
}

func New(opts ...option) (*Con, error) {
	c := Con{
		args:    make(map[string]string),
		context: make(map[string]detail),
	}

	settings := options{
		args: make(map[string]string),
	}

	err := c.resolveOptions(&settings, opts...)
	if err != nil {
		return nil, err
	}

	c.args = settings.args
	c.dir = settings.dir

	err = c.parseDir()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Con) resolveOptions(settings *options, options ...option) error {
	for _, option := range options {
		err := option(settings)
		if err != nil {
			return err
		}
	}

	return nil
}
