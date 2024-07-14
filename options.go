package con

import `strings`

type option func(o *options) error

type options struct {
	dir  string
	args map[string]string
}

func Dir(location string) option {
	return func(o *options) error {
		o.dir = location
		return nil
	}
}

func Arg(key, value string) option {
	return func(o *options) error {
		o.args[strings.ToUpper(key)] = strings.ToUpper(value)
		return nil
	}
}
