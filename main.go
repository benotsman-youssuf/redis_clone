package main

import (
	"flag"
	"log"

	"github.com/benotsman-youssuf/redis_clone/config"
	"github.com/benotsman-youssuf/redis_clone/server"
)

func setupFlags() {
	flag.StringVar(&config.HOST, "host", "0.0.0.0", "host for dice server")
	flag.IntVar(&config.PORT, "port", 7379, "port for dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("rolling the dice 🎲")
	server.RunSyncTCPServer()
}