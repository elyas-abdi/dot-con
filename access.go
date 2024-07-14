package config

import `strings`

func (c *Config) String(key string) *accessStringOption {
	return &accessStringOption{
		accessOption{
			key:       key,
			c:         c,
			accessArg: make(map[string]string),
		},
	}
}

func (o *accessStringOption) Access() *string {
	return o.c.access(o.key)
}

func (c *Config) Slice(key string) *accessStringSliceOption {
	return &accessStringSliceOption{
		accessOption{
			key:       key,
			c:         c,
			accessArg: make(map[string]string),
		},
	}
}

func (o *accessStringSliceOption) Access() *[]string {
	val := o.c.access(o.key)
	if val == nil {
		return nil
	}

	sanitized := *val
	sanitized = sanitized[1 : len(sanitized)-1]
	slice := strings.Split(sanitized, ", ")

	return &slice
}
