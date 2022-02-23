package utils

import "fmt"

func Must(err error) {
	if err != nil {
		panic(fmt.Sprintf("Must: %s\n", err.Error()))
	}
}
