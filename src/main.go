package main

import (
	"fmt"
	"github.com/eremitic/bookstore_users-api/src/app"
	"github.com/spf13/viper"
)

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("default")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	println(123123)

	app.StartApplication()
}
