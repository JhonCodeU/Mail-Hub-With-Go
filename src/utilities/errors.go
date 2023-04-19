package utilities

import "fmt"

func HandleError(err error, message string) {
	if message != "" {
		panic(fmt.Sprintf("%s: %s", message, err))
	} else {
		panic(err)
	}
}
