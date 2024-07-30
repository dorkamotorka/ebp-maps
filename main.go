package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64 maps maps.c

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load pre-compiled programs and maps into the kernel.
	mapsObjs := mapsObjects{}
	if err := loadMapsObjects(&mapsObjs, nil); err != nil {
		log.Fatal(err)
	}
	defer mapsObjs.Close()

	err := mapsObjs.mapsMaps.ArrayMap.Pin("/sys/fs/bpf/test/maps/array_map")
	if err != nil {
		log.Fatal(err)
	}
	defer mapsObjs.mapsMaps.ArrayMap.Unpin()

	err = mapsObjs.mapsMaps.HashMap.Pin("/sys/fs/bpf/test/maps/hash_map")
	if err != nil {
		log.Fatal(err)
	}
	defer mapsObjs.mapsMaps.HashMap.Unpin()

	err = mapsObjs.mapsMaps.LruHashMap.Pin("/sys/fs/bpf/test/maps/lruhash_map")
	if err != nil {
		log.Fatal(err)
	}
	defer mapsObjs.mapsMaps.LruHashMap.Unpin()

	// Set up signal handling to catch Ctrl+C and other interrupts
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	<-sigChan

	log.Println("Interrupt signal received, cleaning up and exiting...")
}

