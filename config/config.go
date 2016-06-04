package configuration

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config is the structure of the
// project's configuration settings
type Config struct {

	/* DB */
	DbName string `default:"go-auth" envconfig:"db_name"`
	DbUser string `default:"root" envconfig:"db_user"`
	DbPass string `required:"true" envconfig:"db_pass"`
	DbHost string `default:"localhost" envconfig:"db_host"`
	// Have to use 32 bit since 16 bit is not enough
	// if the port is greater than 65535
	DbPort int32 `default:"3306" envconfig:"db_port"`

	/* Server */
	Port int32 `default:"8080"`
}

// Export Config for singleton purposes after the
// initialization
var (
	Conf Config
)

// InitConfig is initializing
// the configuration by loading the
// enviroment variables into the project
func InitConfig() Config {
	InitDotenv()
	err := envconfig.Process("goauth", &Conf)
	if err != nil {
		fmt.Println("Error happened while loading enviroment variables for config!")
		log.Fatal(err)
	}
	return Conf
}
