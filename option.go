package config

import `strings`

type accessOption struct {
	key       string
	accessArg map[string]string
	c         *Config
}

type accessStringOption struct{ accessOption }

func (o *accessStringOption) Arg(factor, definition string) *accessStringOption {
	o.accessArg[strings.ToUpper(factor)] = strings.ToUpper(definition)
	return o
}

func (c *Config) String(key string) *accessStringOption {
	return &accessStringOption{
		accessOption{
			key:       key,
			c:         c,
			accessArg: make(map[string]string),
		},
	}
}
