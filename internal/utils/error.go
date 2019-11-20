package utils

import "fmt"

// CheckError Runs panic if the given error isn't nil
func CheckError(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("=> Error:  %s\n   Reason: %s", msg, err))
	}
}
