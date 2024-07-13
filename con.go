package con

type Con struct {
	settings   map[string]string
	properties map[string]interface{}
}

func New() *Con {
	return new(Con)
}

func (c *Con) Property(key string) interface{} {
	value, ok := c.properties[key]
	if !ok {
		return nil
	}

	return value
}
