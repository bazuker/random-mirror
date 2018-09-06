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
	maxStack := 1000
	flag.StringVar(&key, "key", "", "unique key address that accepts new numbers from the remote computer")
	flag.IntVar(&port, "port", 7890, "host port")
	flag.IntVar(&maxStack, "stack", 1000, "maximum stack size")
	flag.Parse()
	if len(key) < 1 {
		panic("key was not provided")
	}
	log.Println("Port:", port)
	log.Println("Key:", key)
	log.Println("Max stack:", maxStack)
	log.Println("Running the mirror ...")
	server := host.NewRandomNumberReceiver(maxStack)
	server.ListenAndServer(":"+strconv.Itoa(port), key)
}
