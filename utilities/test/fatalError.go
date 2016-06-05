package testUtils

import "log"

// CallFatalIfError calls a log.Fatal if error happens
func CallFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
