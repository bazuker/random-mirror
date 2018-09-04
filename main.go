package main

import (
	"log"
	"truerandom-mirror/host"
)

func main() {
	const key = "wslS32bnkb29n1sakSDB3189930SSNssdhH84"
	log.Println("Running the mirror... key is", key)
	server := host.NewRandomNumberReceiver()
	server.ListenAndServer(":7890", key)
}
