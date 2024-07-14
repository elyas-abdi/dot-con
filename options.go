package con

type Option func(o *Options) error

type Options struct {
	dir  string
	args map[string]string
}

func Dir(location string) Option {
	return func(o *Options) error {
		o.dir = location
		return nil
	}
}

func Arg(key, value string) Option {
	return func(o *Options) error {
		o.args[key] = value
		return nil
	}
}
