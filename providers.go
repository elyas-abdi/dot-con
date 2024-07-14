package con

func (c *Con) String(key string) *string {
	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	return &detail.value
}
