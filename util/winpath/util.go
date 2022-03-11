package winpath

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"regexp"
	"strings"
)

func GetRaw(topKey registry.Key, path, value string) (string, error) {
	k, err := registry.OpenKey(topKey, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	s, _, err := k.GetStringValue(value)
	if err != nil {
		return "", err
	}
	return s, nil
}

// GetVal will get a raw value from the registry then replace environment
// variables with their contents.
func GetVal(topKey registry.Key, path, value string) (string, error) {
	v, err := GetRaw(topKey, path, value)
	if err != nil {
		return "", err
	}
	return ReplaceEnv(v)
}

// ReplaceEnv will replace windows environment variables within a given string.
func ReplaceEnv(s string) (string, error) {
	exp := regexp.MustCompile(`%[^%]*%`)
	matches := exp.FindAllString(s, -1)
	for _, m := range matches {
		env := m[1 : len(m)-1]
		if v := os.Getenv(env); v != "" {
			s = strings.Replace(s, m, v, -1)
		}
	}
	return s, nil
}
