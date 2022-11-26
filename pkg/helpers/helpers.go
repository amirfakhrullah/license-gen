package helpers

import (
	"path/filepath"
	"strconv"
	"time"
)

func HandlePanic(e error) {
	if e != nil {
		panic(e)
	}
}

func IsLicenseExist() (bool, error) {
	matches, err := filepath.Glob("LICENSE")
	if err != nil {
		return false, err
	}
	return len(matches) > 0, nil
}

func GetYear() string {
	year := time.Now().Year()
	return strconv.Itoa(year)
}
