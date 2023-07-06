package logger

import "fmt"

func LogError(contextString string, err error) {
	fmt.Printf("%s: %s\n", contextString, err)
}
