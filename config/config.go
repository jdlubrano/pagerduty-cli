package config

import (
  "fmt"
  "os"
  "github.com/spf13/viper"
  homedir "github.com/mitchellh/go-homedir"
)

var configLoaded bool = false

func GetApiToken() string {
  LoadConfig()
  return viper.GetString("api-token")
}

func LoadConfig() {
  if configLoaded {
    return
  }

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".pagerduty")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  configLoaded = true
}
