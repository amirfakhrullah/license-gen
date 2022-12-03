package helpers

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var fileName = "LICENSE"

func HandlePanic(e *error) {
	if *e != nil {
		panic(*e)
	}
}

func IsLicenseExist() ([]string, error) {
	var filesNeededToBeDeleted []string
	for _, licFileName := range []string{"LICENSE", "LICENSE.*"} {
		matches, err := filepath.Glob(licFileName)
		if err != nil {
			return nil, err
		}
		filesNeededToBeDeleted = append(filesNeededToBeDeleted, matches...)

	}
	return filesNeededToBeDeleted, nil
}

func GetDefaultYear() string {
	year := time.Now().Year()
	return strconv.Itoa(year)
}

func DeleteExistingLicenseFiles(existingFiles *[]string) error  {
	for _, lic := range *existingFiles {
		if lic == fileName {
			continue
		}
		err := os.Remove(lic)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateAndWriteLicense(content *string) error {
	f, osErr := os.Create(fileName)
	if osErr != nil {
		return osErr
	}

	_, writeErr := f.WriteString(*content)
	if writeErr != nil {
		return writeErr
	}

	return nil
}