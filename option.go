package config

type accessOption struct {
	key       string
	accessArg map[string]string
	c         *Config
}

type accessStringOption struct{ accessOption }
type accessStringSliceOption struct{ accessOption }
