package exception

func PanicIfNotFoundError(err error) {
	if err != nil {
		panic(NewNotFoundError(err.Error()))
	}
}
