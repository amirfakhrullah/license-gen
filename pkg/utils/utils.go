package utils

func HandlePanic(e error) {
	if e != nil {
		panic(e)
	}
}
