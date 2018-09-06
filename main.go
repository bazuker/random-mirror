package main

import (
	"flag"
	"log"
	"strconv"
	"truerandom-mirror/host"
)

func main() {
	key := ""
	port := 7890
	flag.StringVar(&key, "key", "", "unique key address that accepts new numbers from the remote computer")
	flag.IntVar(&port, "port", 7890, "host port")
	flag.Parse()
	if len(key) < 1 {
		panic("key was not provided")
	}
	log.Println("Port:", port)
	log.Println("Key:", key)
	log.Println("Running the mirror ...")
	server := host.NewRandomNumberReceiver()
	server.ListenAndServer(":"+strconv.Itoa(port), key)
}
