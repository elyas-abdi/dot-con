package config

import `strings`

func (c *Config) Arg(factor, definition string) *Config {
	c.args[strings.ToUpper(factor)] = strings.ToUpper(definition)
	return c
}

func (c *Config) Dir(location string) *Config {
	c.dir = location
	return c
}
