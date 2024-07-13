package con

type Con struct {
	args    map[string]string
	context map[string]interface{}
}

func New(options ...Option) (*Con, error) {
	c := Con{
		args:    make(map[string]string),
		context: make(map[string]interface{}),
	}

	settings := Options{
		file: "",
		dir:  "",
		args: make(map[string]string),
	}

	err := c.resolveOptions(&settings, options...)
	if err != nil {
		return nil, err
	}

	c.args = settings.args

	err = loadEnv()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Con) Context(key string) interface{} {
	value, ok := c.context[key]
	if !ok {
		return nil
	}

	return value
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
