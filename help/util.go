package help

func MyPanic(error interface{}) {
	if error != nil {
		panic(error)
	}
}
