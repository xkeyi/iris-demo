package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

// Viper viper global instance
var Viper *viper.Viper

func init() {
	once.Do(func() {
		Viper = viper.New()
		// scan the file named config in the root directory
		Viper.AddConfigPath("./")
		Viper.SetConfigName("config")

		if err := Viper.ReadInConfig(); err == nil {
			log.Println("Read config successfully: ", Viper.ConfigFileUsed())
		} else {
			log.Printf("Read config failed: %s \n", err)
			panic(err)
		}
	})
}