/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cff/cmd"
)

func initializeViper() {
	/*
		// Set the file name of the configurations file
		viper.SetConfigName("config") // name of config file (without extension)

		// Set the type of the configuration file
		viper.SetConfigType("yaml")

		// Set the path to look for the configurations file
		viper.AddConfigPath(".") // path to look for the config file in

		// Find and read the config file
		err := viper.ReadInConfig()

		if err != nil { // Handle errors reading the config file
			log.Fatalf("Error while reading config file %s", err)
		}*/
}

func main() {
	initializeViper()
	defCmd := "horaire"
	cmd.Execute(defCmd)
}
