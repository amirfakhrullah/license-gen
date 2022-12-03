package helpers

import (
	"path/filepath"
	"strconv"
	"time"
)

func HandlePanic(e *error) {
	if *e != nil {
		panic(*e)
	}
}

func IsLicenseExist() ([]string, error) {
	licenseFileNameList := []string{"LICENSE", "LICENSE.*"}

	var filesNeededToBeDeleted []string
	for _, licFileName := range licenseFileNameList {
		matches, err := filepath.Glob(licFileName)
		if err != nil {
			return nil, err
		}
		filesNeededToBeDeleted = append(filesNeededToBeDeleted, matches...)

	}
	return filesNeededToBeDeleted, nil
}

func GetYear() string {
	year := time.Now().Year()
	return strconv.Itoa(year)
}
