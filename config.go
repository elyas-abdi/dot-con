package config

type Config struct {
	args    map[string]string
	context map[string]detail
	dir     string
}

func New() *Config {
	c := Config{
		args:    make(map[string]string),
		context: make(map[string]detail),
	}

	return &c
}

func (c *Config) Load() (*Config, error) {
	err := c.parseDir()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) access(key string) *string {

	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	return &detail.value
}
