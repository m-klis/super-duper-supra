package exception

import "log"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FailOnError(err error, str string) {
	if err != nil {
		log.Fatalf("Error on : %s", str)
	}
}
