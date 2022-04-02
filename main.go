package main

import (
	"BCC_Intership_BuildUlang/config"
)

func main() {
	_, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
}
