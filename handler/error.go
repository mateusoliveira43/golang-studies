package handler

import (
	"fmt"
)

func HandleError(err error) bool {
	if err != nil {
		fmt.Println("ERROR: An error occurred")
		fmt.Println(err)
		return false
	}
	return true
}
