package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, fmt.Errorf("missing environment variable " + key)
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("environment variable " + key + " should be integer: " + err.Error())
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("environment variable " + key + " should be boolean: " + err.Error())
	}
	return v, nil
}
