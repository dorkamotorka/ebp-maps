package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64 maps maps.c

import (
	"log"
	"time"
)

func main() {
        // Load pre-compiled programs and maps into the kernel.
        mapsObjs := mapsObjects{}
        if err := loadMapsObjects(&mapsObjs, nil); err != nil {
                log.Fatal(err)
        }
        defer mapsObjs.Close()

	for {
		time.Sleep(1 * time.Second)
	}
}
