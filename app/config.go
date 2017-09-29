package app

import (
	"flag"
	"log"
)

type Config struct{

	// Order Engine


	// Validator

}

func (a *App) config(){
	flag.StringVar(&a.ConfigFile, "config", "", "config file to load")
	flag.Parse()

	a.Config = new(Config)
	log.Printf("Loading config file %s..\n", a.ConfigFile)


}

