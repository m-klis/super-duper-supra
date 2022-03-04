package exception

import "log"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
