package exception

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func CheckErrorRespponse(err error) {
	if err != nil {

	}
}
