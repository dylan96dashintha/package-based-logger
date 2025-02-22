package util

import (
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
	"strings"
)

// GetPackageName returns the name of the package where the function is called.
func GetPackageName() (name string, isUnknown bool) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return name, isUnknown
	}

	// Split the file path by slashes
	parts := strings.Split(file, "/")

	// The package name is usually the last part of the file path
	return parts[len(parts)-2], true
}

func YamlReader(filePath string, i interface{}) (err error) {
	byt, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(byt, i)
}
