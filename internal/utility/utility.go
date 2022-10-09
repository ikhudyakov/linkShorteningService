package utility

import (
	"log"
	"os"
)

func GetEnv(text string) string {
	get, ok := os.LookupEnv(text)
	if !ok {
		log.Println("not found environment variable")
		return ""
	}
	return get
}

func CheckError(err error) {
	if err != nil {
		log.Println((err.Error()))
	}
}
