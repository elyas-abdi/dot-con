package config

func (o *accessStringOption) Access() *string {
	return o.c.access(o.key)
}
