package main

import (
	"flag"
	"log"
	"truerandom-mirror/host"
)

func main() {
	key := ""
	flag.StringVar(&key, "key", "", "unique key address that accepts new numbers from the remote computer")
	flag.Parse()
	if len(key) < 1 {
		panic("key was not provided")
	}
	log.Println("Running the mirror... key is", key)
	server := host.NewRandomNumberReceiver()
	server.ListenAndServer(":7890", key)
}
