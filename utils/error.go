package utils

import (
	"fmt"
	"os"
)

func ExitError(err error) {
	fmt.Println(err)
	os.Exit(2)
}