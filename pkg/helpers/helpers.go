package helpers

func HandlePanic(e error) {
	if e != nil {
		panic(e)
	}
}