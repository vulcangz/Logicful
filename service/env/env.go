package env

import "os"

func IsDev() bool {
	return os.Getenv("IS_DEV") == "true"
}
