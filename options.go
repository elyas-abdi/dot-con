package con

type Option func(o *Options) error

type Options struct {
	file string
	dir  string
	args map[string]string
}

func File(location string) Option {
	return func(o *Options) error {
		if len(o.dir) != 0 {
			return ErrDirIsAlreadySpecified(o.dir)
		}

		o.file = location
		return nil
	}
}

func Dir(location string) Option {
	return func(o *Options) error {
		if len(o.file) != 0 {
			return ErrFileIsAlreadySpecified(o.dir)
		}
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
