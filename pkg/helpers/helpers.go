package helpers

import (
	"strconv"
	"time"
)

func HandlePanic(e error) {
	if e != nil {
		panic(e)
	}
}

func GetYear() string {
	year := time.Now().Year()
	return strconv.Itoa(year)
}
