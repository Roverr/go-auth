package configuration

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// InitDotenv is loading the enviroment variables
// from the .env file, which should be located in the
// root of the project
func InitDotenv() {
	var loadError = "open .env: no such file or directory"
	err := godotenv.Load()
	if (err != nil) && (err.Error() != loadError) {
		fmt.Println("Error happened during loading of .env!")
		log.Fatal(err)
	}
}
