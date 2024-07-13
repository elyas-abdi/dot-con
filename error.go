package con

import (
	`errors`
	`fmt`
)

func ErrDirIsAlreadySpecified(dir string) error {
	return errors.New(fmt.Sprintf("specifying a con file is not allowed if a directory is already specified: '%s'", dir))
}

func ErrFileIsAlreadySpecified(file string) error {
	return errors.New(fmt.Sprintf("specifying a con directory is not allowed if file is already specified: '%s'", file))
}

func ErrUnableToLoadDotEnvFile() error {
	return errors.New("unable to load .env file")
}
