// package Config get environment settings
package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// config structure
type config struct {
	Server struct {
		Address string
		Port    int
		Timeout time.Duration
	}
	CsvPath struct {
		Prod string
		Test string
	}
	Api struct {
		Url string
	}
}

// C config global var type config
var C config

// ReadConfig read YML file convert to config struct
func ReadConfig(configFile string) error {
	Config := &C
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}