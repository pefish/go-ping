package main

import (
	"flag"
	"fmt"
	go_ping "github.com/pefish/go-ping"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go-ping host\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	addr := flag.Arg(0)
	dst, dur, err := go_ping.Ping(addr)
	if err != nil {
		log.Fatalf("Ping %s (%s): %s\n", addr, dst, err)
	}
	fmt.Printf("Ping %s (%s): %s\n", addr, dst, dur)
}
