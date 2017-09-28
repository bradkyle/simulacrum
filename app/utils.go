package app

import (
	"os"
	"runtime"
	"strconv"
	"os/signal"
	"syscall"
	"log"
	"net/http"
	"encoding/json"
)


// AdjustGoMaxProcs adjusts the maximum processes that the CPU can handle.
func (a *App) AdjustGoMaxProcs() {
	log.Println("Adjusting bot runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	log.Println("Number of CPU's detected:", maxProcs)

	if maxProcsEnv != "" {
		log.Println("GOMAXPROCS env =", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)
		if err != nil {
			log.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
		} else {
			maxProcs = env
		}
	}
	if i := runtime.GOMAXPROCS(maxProcs); i != maxProcs {
		log.Fatal("Go Max Procs were not set correctly.")
	}
	log.Println("Set GOMAXPROCS to:", maxProcs)
}


// HandleInterrupt monitors and captures the SIGTERM in a new goroutine then
// shuts down bot
func (a *App) HandleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Captured %v.", sig)
		a.Shutdown()
	}()
}

// Shutdown correctly shuts down bot saving configuration files
func (a *App) Shutdown() {
	<-a.shutdown
	log.Println("Bot shutting down..")

	// do stuff before shutdown

	log.Println("Exiting.")
	os.Exit(1)
}
