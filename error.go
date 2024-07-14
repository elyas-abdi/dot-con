package config

import (
	`errors`
	`fmt`
)

func errDirIsAlreadySpecified(dir string) error {
	return errors.New(fmt.Sprintf("specifying a con file is not allowed if a directory is already specified: '%s'", dir))
}

func errFileIsAlreadySpecified(file string) error {
	return errors.New(fmt.Sprintf("specifying a con directory is not allowed if file is already specified: '%s'", file))
}

func errUnableToLoadDotEnvFile() error {
	return errors.New("unable to load .env file")
}
