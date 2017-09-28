package app

import (
	"flag"
	"log"
)

type Config struct{



	// Validator

}

func (a *App) loadConfig(){
	flag.StringVar(&a.ConfigFile, "config", "", "config file to load")
	flag.Parse()

	a.Config = new(Config)
	log.Printf("Loading config file %s..\n", a.ConfigFile)


}

