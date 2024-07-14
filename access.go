package config

import (
	`strconv`
	`strings`
)

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

func (c *Config) Int(key string) *accessInt64Option {
	return &accessInt64Option{
		accessOption{
			key:       key,
			c:         c,
			accessArg: make(map[string]string),
		},
	}
}

func (o *accessInt64Option) Access() *int64 {
	val := o.c.access(o.key)
	if val == nil {
		return nil
	}

	integer, err := strconv.ParseInt(*val, 0, 64)
	if err != nil {
		return nil
	}

	return &integer
}
