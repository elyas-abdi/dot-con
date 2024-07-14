package con

import `strconv`

func (c *Con) String(key string) *string {
	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	return &detail.value
}

func (c *Con) Bool(key string) *bool {
	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	val, err := strconv.ParseBool(detail.value)
	if err != nil {
		return nil
	}

	return &val
}

func (c *Con) Int(key string) *int64 {
	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	val, err := strconv.ParseInt(detail.value, 0, 64)
	if err != nil {
		return nil
	}

	return &val
}

func (c *Con) Float(key string) *float64 {
	detail, ok := c.context[key]
	if !ok {
		return nil
	}

	val, err := strconv.ParseFloat(detail.value, 64)
	if err != nil {
		return nil
	}

	return &val
}
