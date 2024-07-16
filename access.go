package config

import (
	"strconv"
	"strings"
)

func (c *Config) Access(key string) *string {
	return c.access(key)
}

func (c *Config) AccessInt(key string) *int64 {
	val := c.access(key)
	if val == nil {
		return nil
	}

	integer, err := strconv.ParseInt(*val, 0, 64)
	if err != nil {
		return nil
	}

	return &integer
}

func (c *Config) AccessSlice(key string) *[]string {
	val := c.access(key)
	if val == nil {
		return nil
	}

	sanitized := *val
	sanitized = sanitized[1 : len(sanitized)-1]
	slice := strings.Split(sanitized, ", ")

	return &slice
}
